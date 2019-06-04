package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/iwaltgen/grpc-gogo-test/pkg/gogoserver/grpc/service"
	pb "github.com/iwaltgen/grpc-gogo-test/proto"
	"google.golang.org/grpc"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// TODO(iwaltgen): flags
	addr := ":8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, service.New())

	fmt.Printf("Serving gRPC %v", addr)
	if err := s.Serve(ln); err != nil {
		fmt.Fprintf(os.Stdout, "error serve: %v\n", err)
	}
}
