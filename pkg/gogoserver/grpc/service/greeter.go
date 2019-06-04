package service

import (
	"context"
	"errors"
	"time"

	"github.com/gogo/protobuf/types"
	pb "github.com/iwaltgen/grpc-gogo-test/proto/v1"
)

// Greeter implement grpc GreeterServiceServer
type Greeter struct {
}

// New Greeter
func New() *Greeter {
	return &Greeter{}
}

// SayHello name hello reply
func (s *Greeter) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: req.Name + " hello",
	}, nil
}

// SayTime current time reply
func (s *Greeter) SayTime(ctx context.Context, req *types.Empty) (*types.Timestamp, error) {
	now := time.Now()
	return &types.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.Nanosecond()),
	}, nil
}

// SubscribeTime multiple current time reply
func (s *Greeter) SubscribeTime(
	req *pb.TimeRequest,
	stream pb.GreeterService_SubscribeTimeServer) error {

	ctx := stream.Context()
	interval, _ := types.DurationFromProto(req.Interval)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for i := int32(0); i < req.Count; i++ {
		select {
		case <-ctx.Done():
			return errors.New("client cancel")

		case now := <-ticker.C:
			ts, err := types.TimestampProto(now)
			if err != nil {
				return err
			}
			if err := stream.Send(ts); err != nil {
				return err
			}
		}
	}

	return nil
}
