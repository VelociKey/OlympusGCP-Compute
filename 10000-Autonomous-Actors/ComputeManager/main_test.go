package main

import (
	"context"
	"testing"

	computev1 "OlympusGCP-Compute/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/compute/v1"
	"connectrpc.com/connect"
)

func TestComputeServerAdvanced(t *testing.T) {
	server := &ComputeServer{}
	ctx := context.Background()

	// Test CheckHealth
	req := connect.NewRequest(&computev1.CheckHealthRequest{ServiceName: "api"})
	res, err := server.CheckHealth(ctx, req)
	if err != nil {
		t.Fatalf("CheckHealth failed: %v", err)
	}
	if res.Msg.Status != computev1.CheckHealthResponse_HEALTHY {
		t.Errorf("Expected HEALTHY status, got %v", res.Msg.Status)
	}
}
