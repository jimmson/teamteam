package main

import (
	"flag"
	"net/http"

	"github.com/corverroos/unsure"
	teamteam_ops "github.com/jimmson/teamteam/ops"
	teamteam_server "github.com/jimmson/teamteam/server"
	"github.com/jimmson/teamteam/state"
	"github.com/jimmson/teamteam/teamteampb"
	"github.com/luno/jettison/errors"
)

var (
	httpAddress = flag.String("http_address", ":23047", "teamteam healthcheck address")
	grpcAddress = flag.String("grpc_address", ":23048", "teamteam grpc server address")
)

func main() {
	unsure.Bootstrap()

	s, err := state.New()
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new state error"))
	}

	go serveGRPCForever(s)

	teamteam_ops.StartLoops(s)

	http.HandleFunc("/health", makeHealthCheckHandler())
	go unsure.ListenAndServeForever(*httpAddress, nil)

	unsure.WaitForShutdown()
}

func serveGRPCForever(s *state.State) {
	grpcServer, err := unsure.NewServer(*grpcAddress)
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new grpctls server"))
	}

	teamteamSrv := teamteam_server.New(s)
	teamteampb.RegisterTeamTeamServer(grpcServer.GRPCServer(), teamteamSrv)

	unsure.RegisterNoErr(func() {
		//teamteamSrv.Stop()
		grpcServer.Stop()
	})

	unsure.Fatal(grpcServer.ServeForever())
}

func makeHealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}
}
