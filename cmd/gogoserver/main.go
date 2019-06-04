package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/iwaltgen/grpc-gogo-test/pkg/gogoserver/grpc"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// TODO(iwaltgen): flags
	if err := grpc.Serve(8080); err != nil {
		fmt.Fprintln(os.Stderr, "grpc serve error: %v", err)
	}
}
