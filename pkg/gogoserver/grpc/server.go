package grpc

import (
	"fmt"
	"net"

	"github.com/iwaltgen/grpc-gogo-test/pkg/gogoserver/grpc/service"
	pb "github.com/iwaltgen/grpc-gogo-test/proto/v1"
	"google.golang.org/grpc"
)

// Serve gRPC Services
func Serve(port int) error {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, service.New())

	fmt.Printf("Serving gRPC %v", addr)
	if err := s.Serve(ln); err != nil {
		return fmt.Errorf("error serve: %v", err)
	}
	return nil
}
