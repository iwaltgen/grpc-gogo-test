package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"
	pb "github.com/iwaltgen/grpc-gogo-test/proto/v1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// greeterCmd represents the greeter command
var greeterCmd = &cobra.Command{
	Use:   "greeter",
	Short: "Request greeterService",
	Long: `Request greeterService.

'SayHello john' reply john hello
'SayTime' reply current time
'SubscribeTime 3 1s' reply current time 3 times interval 1s

Cobra application test project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("greeter length args must greater than 1")
		}

		rpcMethod := strings.ToLower(args[0])
		switch rpcMethod {
		case "sayhello":
			if len(args) < 2 {
				return errors.New("greeter SayHello length args must greater than 2")
			}

			return command(sayHello(args[1]))

		case "saytime":
			return command(sayTime())

		case "subscribetime":
			if len(args) < 3 {
				return errors.New("greeter SubscribeTime length args must greater than 3")
			}

			count, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			interval, err := time.ParseDuration(args[2])
			if err != nil {
				return err
			}
			return command(subscribeTime(count, interval))

		default:
			return fmt.Errorf("unexpected rpc method name: %s", rpcMethod)
		}
	},
}

func init() {
	rootCmd.AddCommand(greeterCmd)
}

func newClient() (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addr := "127.0.0.1:8080"
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s [%v]", addr, err)
	}
	return conn, nil
}

type procedure func(pb.GreeterServiceClient) error

func command(cmd procedure) error {
	conn, err := newClient()
	if err != nil {
		return err
	}
	defer conn.Close()

	return cmd(pb.NewGreeterServiceClient(conn))
}

func sayHello(name string) procedure {
	return func(c pb.GreeterServiceClient) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		reply, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.FailFast(true))
		if err != nil {
			return fmt.Errorf("rpc error: %v", err)
		}

		fmt.Println(reply)
		return nil
	}
}

func sayTime() procedure {
	return func(c pb.GreeterServiceClient) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		reply, err := c.SayTime(ctx, &types.Empty{}, grpc.FailFast(true))
		if err != nil {
			return fmt.Errorf("rpc error: %v", err)
		}

		fmt.Println(reply)
		return nil
	}
}

func subscribeTime(count int, interval time.Duration) procedure {
	return func(c pb.GreeterServiceClient) error {
		timeout := time.Duration(count+1) * interval
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		req := &pb.TimeRequest{
			Count:    int32(count),
			Interval: types.DurationProto(interval),
		}
		subscriber, err := c.SubscribeTime(ctx, req, grpc.FailFast(true))
		if err != nil {
			return fmt.Errorf("rpc error: %v", err)
		}

		for {
			reply, err := subscriber.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("subscribe error: %v", err)
			}
			fmt.Println(reply)
		}
		return nil
	}
}
