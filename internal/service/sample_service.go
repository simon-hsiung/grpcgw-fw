package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"main/internal/utils"
	"main/protocol/pb/v1"
	"math/rand"
	"strconv"

	"github.com/TXOne-Stellar/stellar-lib/errutil/richerr"
	"github.com/TXOne-Stellar/stellar-lib/middleware/grpcmw"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

func NewSampleService() *SampleService {
	return &SampleService{}
}

type SampleService struct {
	pb.UnimplementedSampleServiceServer
}

func (s *SampleService) RetrieveSampleData(ctx context.Context, req *pb.SampleRequest) (
	*pb.SampleResponse, error,
) {
	msg := req.GetGuid()
	if msg == "" {
		msg = strconv.Itoa(int(req.GetId()))
	}
	return &pb.SampleResponse{Data: msg}, nil
}

func (s *SampleService) StreamDownload(
	req *pb.StreamRequest, stream pb.SampleService_StreamDownloadServer,
) error {
	source := rand.NewSource(req.GetSeed())
	r := rand.New(source)
	for i := int32(0); i < req.GetIteration(); i++ {
		resp := &pb.StreamResponse{Payload: make([]byte, req.SectionLength)}
		for i := 0; i < int(req.SectionLength); i++ {
			resp.Payload[i] = byte(r.Intn(256))
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

func (s *SampleService) StreamDownloadHttp(
	req *pb.StreamRequest, stream pb.SampleService_StreamDownloadHttpServer,
) error {
	source := rand.NewSource(req.GetSeed())
	r := rand.New(source)
	for i := int32(0); i < req.GetIteration(); i++ {
		data := make([]byte, req.SectionLength)
		for i := 0; i < int(req.SectionLength); i++ {
			data[i] = byte(r.Intn(256))
		}
		if err := stream.Send(&httpbody.HttpBody{
			ContentType: "application/octet-stream",
			Data:        data,
		}); err != nil {
			return utils.ErrInternal().SetUnderlying(err).SetMessage("failed to send data")
		}
	}

	return nil
}

func (s *SampleService) StreamUploadHttp(srv pb.SampleService_StreamUploadHttpServer) error {
	var reader io.ReadCloser

	// get multipart form
	if form, err := grpcmw.NewMultipartStreamReader(srv).ReadForm(); err == nil {
		// get uploaded file descriptor
		file, err := form.OpenFile("data_file")
		if err != nil {
			return utils.ErrArgInvalid().SetUnderlying(err).SetMessage("cannot open data_file")
		}
		reader = file
	} else {
		// if it's header error, it may be passed through normal grpc
		if richerr.IsCausedBy(grpcmw.ErrCodeIncorrectHeader.Code(), err) {
			// try to use http body reader
			reader = grpcmw.NewHttpBodyStreamReader(srv)
		} else {
			return err
		}
	}

	defer reader.Close()

	buff := make([]byte, 1024)
	hasher := sha256.New()
	for {
		n, err := reader.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			return utils.ErrInternal().SetUnderlying(err).SetMessage("failed to read data")
		}
		hasher.Write(buff[:n])
	}

	hash := hasher.Sum(nil)
	return srv.SendAndClose(&pb.SampleResponse{Data: hex.EncodeToString(hash)})
}
