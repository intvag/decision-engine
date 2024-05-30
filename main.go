package main

import (
	"net"
	"os"
	"strconv"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/intvag/decision-engine/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	ProfitabilityMultiplier = envWithDefaultFloat("DECISION_ENGINE_PROFITABILITY", 1.1)
	ListenAddr              = envOrDefaultString("DECISION_ENGINE_LISTEN_ADDR", "0.0.0.0:8888")
)

func main() {
	// #nosec: G102
	lis, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		panic(err)
	}

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(InterceptorLogger(), loggingOpts...),
	))

	reflection.Register(grpcServer)
	service.RegisterQuotesServer(grpcServer, new(Server))

	panic(grpcServer.Serve(lis))
}

func envWithDefaultFloat(k string, d float64) float64 {
	v, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 1
	}

	return f
}

func envOrDefaultString(k, d string) string {
	v, ok := os.LookupEnv(k)
	if ok {
		return v
	}

	return d
}
