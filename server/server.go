package server

import (
	"context"

	"github.com/jimmson/teamteam/db/events"
	pb "github.com/jimmson/teamteam/teamteampb"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
)

var _ pb.TeamTeamServer = (*Server)(nil)

// Server implements the engine grpc server.
type Server struct {
	b       Backends
	rserver *reflex.Server
	stream  reflex.StreamFunc
}

// New returns a new server instance.
func New(b Backends) *Server {
	return &Server{
		b:       b,
		rserver: reflex.NewServer(),
		stream:  events.ToStream(b.TeamteamDB().DB),
	}
}

func (srv *Server) Stop() {
	srv.rserver.Stop()
}

func (srv *Server) Ping(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return req, nil
}

func (srv *Server) Stream(req *reflexpb.StreamRequest, ss pb.TeamTeam_StreamServer) error {
	return srv.rserver.Stream(srv.stream, req, ss)
}

func (srv *Server) GetRound(context.Context, *pb.GetRoundRequest) (*pb.Round, error) {
	return &pb.Round{
		PlayerRank: 3,
		MyPart:     2,
	}, nil
}
