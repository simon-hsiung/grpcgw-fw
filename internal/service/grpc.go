package service

import (
	"context"
	"main/internal/utils"
	"main/protocol/pb/v1"
	"net"
	"runtime/debug"

	"github.com/TXOne-Stellar/stellar-lib/logging"
	"github.com/TXOne-Stellar/stellar-lib/middleware/grpcmw"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Address      string
	Server       *grpc.Server
	StartLogging func()
	StopLogging  func()
}

func (s *GrpcServer) Start() error {
	lis, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	defer lis.Close()

	if s.StartLogging != nil {
		s.StartLogging()
	}
	return s.Server.Serve(lis)
}

func (s *GrpcServer) Stop() error {
	s.Server.GracefulStop()
	if s.StopLogging != nil {
		s.StopLogging()
	}
	return nil
}

func MajorGrpcServer(addr string) *GrpcServer {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpcmw.PeerIp().UnaryServerInterceptor(),
				grpcmw.RequestId().UnaryServerInterceptor(),
				grpcmw.Logging(logging.Default()).UnaryServerInterceptor(),
				grpcmw.Validating().UnaryServerInterceptor(),
				grpc_recovery.UnaryServerInterceptor(
					grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p any) error {
						logging.WithCtx(ctx).Error("panic: %v\nstacktrace: %s", p, string(debug.Stack()))
						return utils.ErrInternal().SetMessage("server panic: %v", p)
					}),
				),
			),
		),
		grpc.ChainStreamInterceptor(
			grpcmw.PeerIp().StreamServerInterceptor(),
			grpcmw.RequestId().StreamServerInterceptor(),
			grpcmw.Logging(logging.Default()).StreamServerInterceptor(),
			grpcmw.Validating().StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(
				grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p any) (err error) {
					logging.WithCtx(ctx).Error("panic: %v\nstacktrace: %s", p, string(debug.Stack()))
					return utils.ErrInternal().SetMessage("server panic: %v", p)
				}),
			),
		),
	)

	pb.RegisterSampleServiceServer(s, NewSampleService())

	return &GrpcServer{
		Address:      addr,
		Server:       s,
		StartLogging: func() { logging.Info("Serving gRPC at %s", addr) },
		StopLogging:  func() { logging.Info("gRPC service stopped gracefully") },
	}
}
