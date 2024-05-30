package main

import (
	"net"
	"os"
	"strconv"

	"github.com/intvag/decision-engine/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	ProfitabilityMultiplier = envWithDefaultFloat("DECISION_ENGINE_PROFITABILITY", 1.1)
)

func main() {
	// #nosec: G102
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	service.RegisterQuotesServer(grpcServer, new(Server))

	panic(grpcServer.Serve(lis))
}

func envWithDefaultFloat(k string, d float64) float64 {
	v, ok := os.LookupEnv("DECISION_ENGINE_PROFITABILITY")
	if !ok {
		return d
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 1
	}

	return f
}
