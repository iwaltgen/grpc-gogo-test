package main

import (
	"math/rand"
	"time"

	"github.com/iwaltgen/grpc-gogo-test/pkg/gogoctl/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
