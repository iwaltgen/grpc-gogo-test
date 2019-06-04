package grpc

import (
	"fmt"
	"net"
	"os"
	"runtime"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/iwaltgen/grpc-gogo-test/pkg/gogoserver/grpc/service"
	pb "github.com/iwaltgen/grpc-gogo-test/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Serve gRPC Services
func Serve(port int) error {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
	)
	pb.RegisterGreeterServiceServer(s, service.New())

	fmt.Printf("Serving gRPC %v", addr)
	if err := s.Serve(ln); err != nil {
		return fmt.Errorf("error serve: %v", err)
	}
	return nil
}

var panicHandler = grpc_recovery.RecoveryHandlerFunc(func(p interface{}) error {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	fmt.Fprintf(os.Stderr, "panic recovered: %+v", string(buf))
	return status.Errorf(codes.Internal, "%s", p)
})
