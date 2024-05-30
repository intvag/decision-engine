package main

import (
	"context"
	"math"

	"github.com/intvag/decision-engine/service"
)

type Server struct {
	service.UnimplementedQuotesServer
}

func (s Server) GetQuote(ctx context.Context, in *service.Input) (q *service.Quote, err error) {
	raw := (in.ExpectedLifetime - in.Age) / 12 / in.Lastability * in.Repairability * ProfitabilityMultiplier

	return &service.Quote{
		Monthly: math.Ceil(raw*100) / 100,
	}, nil
}
