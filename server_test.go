package main

import (
	"context"
	"testing"

	"github.com/intvag/decision-engine/service"
)

func TestServer_GetQuote(t *testing.T) {
	s := new(Server)

	q, err := s.GetQuote(context.Background(), &service.Input{
		Age:              0,
		PurchasePrice:    900,
		ExpectedLifetime: 20,
		Lastability:      0.45,
		Repairability:    0.75,
	})
	if err != nil {
		t.Fatal(err)
	}

	expect := 3.06
	if expect != q.Monthly {
		t.Errorf("expected %v, received %v", expect, q.Monthly)
	}
}
