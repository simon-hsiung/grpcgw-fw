package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/internal/utils"
	"main/protocol/pb/v1"
	"net/http"
	"time"

	"github.com/TXOne-Stellar/stellar-lib/logging"
	"github.com/TXOne-Stellar/stellar-lib/middleware/grpcmw"
	"github.com/TXOne-Stellar/stellar-lib/middleware/httpmw"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/encoding/protojson"
)

type HttpServer struct {
	CertFilePath string
	KeyFilePath  string
	Server       *http.Server
	StartLogger  func()
	StopLogger   func()
}

func (s *HttpServer) Start() error {
	if s.StartLogger != nil {
		s.StartLogger()
	}

	err := s.Server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return utils.ErrFileIO().SetMessage("failed to serve at %s: %v", s.Server.Addr, err)
	}

	return nil
}

func (s *HttpServer) Stop() error {
	ctxShutdown, abort := context.WithTimeout(context.Background(), 10*time.Second)
	defer abort()

	if err := s.Server.Shutdown(ctxShutdown); err != nil {
		return utils.ErrInternal().SetMessage("failed to shutdown http server: %v", err)
	}

	if s.StopLogger != nil {
		s.StopLogger()
	}

	return nil
}

func MajorHttpServer(ctx context.Context, addr string, grpcConn *grpc.ClientConn) *HttpServer {
	// grpc gateway mux
	mux := runtime.NewServeMux(
		// allow incoming custom headers
		runtime.WithIncomingHeaderMatcher(grpcmw.CustomHeaderMatcher),
		// allow all outgoing headers
		runtime.WithOutgoingHeaderMatcher(grpcmw.NoopHeaderMatcher),
		// accept raw protobuf for demo usage
		runtime.WithMarshalerOption("application/protobuf", &runtime.ProtoMarshaller{}),
		// default marshaler to output json or http body
		runtime.WithMarshalerOption(runtime.MIMEWildcard, grpcmw.NewHttpBodyMarshaler()),
		// error handler
		runtime.WithErrorHandler(grpcmw.Handlers().RpcErrorHandlerV2()),
		// health check
		runtime.WithHealthEndpointAt(grpc_health_v1.NewHealthClient(grpcConn), "/healthz"),
	)

	// register web grpc service handlers to grpc gateway
	if err := pb.RegisterSampleServiceHandler(ctx, mux, grpcConn); err != nil {
		panic(fmt.Sprintf("failed to register web grpc-gateway handler: %v", err))
	}

	// special handling for stream download
	g_grpcConn = grpcConn
	if err := mux.HandlePath(
		http.MethodPost,
		"/v1/sample/stream/{seed}/with-handler",
		streamDownloadHandler,
	); err != nil {
		logging.Fatal(err.Error())
	}

	// bind logging middleware
	handler := httpmw.Logging(logging.Default()).Handler(mux)
	// bind request id middleware
	handler = httpmw.RequestId().Handler(handler)

	// http server
	httpMux := http.NewServeMux()
	httpMux.Handle("/", handler)
	// skip logging for health check
	httpMux.Handle("/healthz", mux)

	srv := http.Server{
		Addr:    addr,
		Handler: httpMux,
	}

	return &HttpServer{
		Server:      &srv,
		StartLogger: func() { logging.Info("Serving HTTP on: %s", srv.Addr) },
		StopLogger:  func() { logging.Info("HTTP server stopped gracefully") },
	}
}

var g_grpcConn *grpc.ClientConn

func streamDownloadHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	req, err := buildGrpcRequest[pb.StreamRequest](r, params, []string{"seed"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	stream, err := pb.NewSampleServiceClient(g_grpcConn).StreamDownload(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(resp.Payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func buildGrpcRequest[RequestType any](
	r *http.Request, params map[string]string, necessaryFields []string,
) (
	*RequestType, error,
) {
	// read request body
	buffer := bytes.Buffer{}
	if _, err := io.Copy(&buffer, r.Body); err != nil {
		return nil, utils.ErrInternal().SetUnderlying(err).SetMessage("read request body failed")
	}

	// re-compose the json body to include necessary fields passed in params map if any
	body := buffer.Bytes()
	if (len(necessaryFields)) > 0 && len(params) > 0 {
		// unmarshal the json body
		var jsonBody map[string]any
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			return nil, utils.ErrArgInvalid().SetUnderlying(err).SetMessage("invalid json body")
		}
		// add necessary fields to the json body
		for _, key := range necessaryFields {
			jsonBody[key] = params[key]
		}

		// re-marshal json to bytes
		var err error
		body, err = json.Marshal(jsonBody)
		if err != nil {
			return nil, utils.ErrMarshaling().SetUnderlying(err).
				SetMessage("re-marshal request body failed")
		}
	}

	// unmarshal the json bytes to the request type
	unmarshaler := runtime.JSONPb{
		UnmarshalOptions: protojson.UnmarshalOptions{
			AllowPartial:   true,
			DiscardUnknown: true,
		},
	}
	var out RequestType
	if err := unmarshaler.Unmarshal(body, &out); err != nil {
		return nil, utils.ErrMarshaling().SetUnderlying(err).
			SetMessage("unmarshal request body failed")
	}

	return &out, nil
}
