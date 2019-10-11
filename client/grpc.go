package client

import (
	"context"
	"flag"

	"github.com/jimmson/teamteam"
	pb "github.com/jimmson/teamteam/teamteampb"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
	"google.golang.org/grpc"
)

var addr = flag.String("teamteam_addresses", "", "host:port of engine gRPC service")

type client struct {
	address   string
	rpcConn   *grpc.ClientConn
	rpcClient pb.TeamTeamClient
}

func (c *client) Ping(ctx context.Context) error {
	_, err := c.rpcClient.Ping(ctx, &pb.Empty{})
	return err
}

func (c *client) Stream(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error) {
	sFn := reflex.WrapStreamPB(func(ctx context.Context,
		req *reflexpb.StreamRequest) (reflex.StreamClientPB, error) {
		return c.rpcClient.Stream(ctx, req)
	})
	return sFn(ctx, after, opts...)
}

func (c *client) GetRound(ctx context.Context, roundID int, playerName string) (*teamteam.Round, error) {
	res, err := c.rpcClient.GetRound(ctx, &pb.GetRoundRequest{
		Round:  int32(roundID),
		Player: playerName,
	})
	if err != nil {
		return nil, err
	}
	return &teamteam.Round{
		PlayerRank: res.PlayerRank,
		MyPart:     res.MyPart,
	}, err
}
