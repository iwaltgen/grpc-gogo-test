package main

import (
	"math/rand"
	"time"

	"github.com/iwaltgen/grpc-gogo-test/pkg/gogoctl/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
