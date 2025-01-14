package service

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"main/protocol/pb/v1"
	"math/rand"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestSampleService(t *testing.T) {
	suite.Run(t, &SampleServiceTestSuite{
		GrpcAddr: "localhost:9999",
		HttpAddr: "localhost:8888",
	})

}

type SampleServiceTestSuite struct {
	suite.Suite

	GrpcAddr string
	HttpAddr string

	grpcSrv  *GrpcServer
	grpcConn *grpc.ClientConn
	ctx      context.Context
	finish   context.CancelFunc
	httpSrv  *HttpServer
}

func (s *SampleServiceTestSuite) SetupSuite() {
	s.grpcSrv = MajorGrpcServer(s.GrpcAddr)
	go func() {
		if err := s.grpcSrv.Start(); err != nil {
			s.FailNow("failed to start grpc server: %v", err)
		}
	}()

	var err error
	s.grpcConn, err = grpc.NewClient(
		s.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		s.FailNow("failed to create grpc connection: %v", err)
	}

	s.ctx, s.finish = context.WithCancel(context.Background())

	s.httpSrv = MajorHttpServer(s.ctx, s.HttpAddr, s.grpcConn)
	go func() {
		if err := s.httpSrv.Start(); err != nil {
			s.FailNow("failed to start http server: %v", err)
		}
	}()
}

func (s *SampleServiceTestSuite) TearDownSuite() {
	s.Assert().Nil(s.httpSrv.Stop())
	s.finish()
	s.Assert().Nil(s.grpcConn.Close())
	s.Assert().Nil(s.grpcSrv.Stop())
}

func (s *SampleServiceTestSuite) TestRetrieveSampleData() {
	req := &pb.SampleRequest{Guid: uuid.New().String()}
	jsb, _ := json.Marshal(req)
	addr := fmt.Sprintf("http://%s/v1/sample/data", s.HttpAddr)
	ret, err := http.Post(addr, "application/json", bytes.NewReader(jsb))
	s.Require().Nil(err)
	jsb, _ = io.ReadAll(ret.Body)
	var resp pb.SampleResponse
	s.Require().Nil(json.Unmarshal(jsb, &resp))

	s.Require().Equal(req.Guid, resp.Data)
}

func randomData(sectionLen int32) (int64, int32, []byte) {
	seed := rand.Int63n(100)
	iteration := rand.Int31() % 10

	source := rand.NewSource(seed)
	r := rand.New(source)

	dataLen := sectionLen * iteration
	data := make([]byte, dataLen)
	for i := int32(0); i < dataLen; i++ {
		data[i] = byte(r.Intn(256))
	}

	return seed, iteration, data
}

func (s *SampleServiceTestSuite) TestStreamDownload() {
	sectionLen := int32(10)
	seed, iteration, data := randomData(sectionLen)

	// grpc
	stream, err := pb.NewSampleServiceClient(s.grpcConn).StreamDownload(
		s.ctx,
		&pb.StreamRequest{
			Seed:          seed,
			Iteration:     iteration,
			SectionLength: sectionLen,
		},
	)
	s.Require().Nil(err)
	for it := int32(0); it < iteration; it++ {
		resp, err := stream.Recv()
		s.Require().Nil(err)
		for i := int32(0); i < sectionLen; i++ {
			pos := it*sectionLen + i
			s.Require().Equal(data[pos], resp.Payload[i])
		}
	}

	// http
	jsb, _ := json.Marshal(&pb.StreamRequest{Iteration: iteration, SectionLength: sectionLen})
	addr := fmt.Sprintf("http://%s/v1/sample/stream/%d", s.HttpAddr, seed)
	ret, err := http.Post(addr, "application/json", bytes.NewReader(jsb))
	s.Require().Nil(err)
	retBytes, _ := io.ReadAll(ret.Body)
	s.Require().NotEqual(data, retBytes)

	// http with handler
	addr = fmt.Sprintf("http://%s/v1/sample/stream/%d/with-handler", s.HttpAddr, seed)
	ret, err = http.Post(addr, "application/json", bytes.NewReader(jsb))
	s.Require().Nil(err)
	retBytes, _ = io.ReadAll(ret.Body)
	s.Require().Equal(data, retBytes)
}

func (s *SampleServiceTestSuite) TestStreamDownloadHttp() {
	sectionLen := int32(10)
	seed, iteration, data := randomData(sectionLen)

	// grpc
	stream, err := pb.NewSampleServiceClient(s.grpcConn).StreamDownloadHttp(
		s.ctx,
		&pb.StreamRequest{
			Seed:          seed,
			Iteration:     iteration,
			SectionLength: sectionLen,
		},
	)
	s.Require().Nil(err)
	for it := int32(0); it < iteration; it++ {
		resp, err := stream.Recv()
		s.Require().Nil(err)
		for i := int32(0); i < sectionLen; i++ {
			pos := it*sectionLen + i
			s.Require().Equal(data[pos], resp.Data[i])
		}
	}

	// http
	jsb, _ := json.Marshal(&pb.StreamRequest{Iteration: iteration, SectionLength: sectionLen})
	addr := fmt.Sprintf("http://%s/v1/sample/stream/http/%d", s.HttpAddr, seed)
	ret, err := http.Post(addr, "application/json", bytes.NewReader(jsb))
	s.Require().Nil(err)
	retBytes, _ := io.ReadAll(ret.Body)
	s.Require().Equal(data, retBytes)
}

func (s *SampleServiceTestSuite) TestStreamDownloadHttpHandler() {
	sectionLen := int32(100)
	_, iteration, data := randomData(sectionLen)
	hashBytes := sha256.Sum256(data)
	hash := hex.EncodeToString(hashBytes[:])

	// grpc
	stream, err := pb.NewSampleServiceClient(s.grpcConn).StreamUploadHttp(s.ctx)
	s.Require().Nil(err)
	for it := int32(0); it < iteration; it++ {
		pos := it * sectionLen
		err := stream.Send(&httpbody.HttpBody{
			ContentType: "application/octet-stream",
			Data:        data[pos : pos+sectionLen],
		})
		s.Require().True(err == nil || err == io.EOF, "failed to send data: %v", err)
	}
	resp, err := stream.CloseAndRecv()
	s.Require().Nil(err, "failed to receive response: %v", err)
	s.Require().Equal(hash, resp.Data)

	// http
	addr := fmt.Sprintf("http://%s/v1/sample/stream/upload/http", s.HttpAddr)
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("data_file", "file_name")
	s.Require().Nil(err)
	_, _ = io.Copy(fw, bytes.NewReader(data))
	writer.Close()
	ret, err := http.Post(addr, writer.FormDataContentType(), body)
	s.Require().Nil(err)
	jsb, _ := io.ReadAll(ret.Body)
	var httpResp pb.SampleResponse
	err = json.Unmarshal(jsb, &httpResp)
	s.Require().Nil(err)
	s.Require().Equal(hash, httpResp.Data)
}
