package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	computev1 "OlympusGCP-Compute/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/compute/v1"
	"OlympusGCP-Compute/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/compute/v1/computev1connect"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ComputeServer struct{}

func (s *ComputeServer) RunService(ctx context.Context, req *connect.Request[computev1.RunServiceRequest]) (*connect.Response[computev1.RunServiceResponse], error) {
	slog.Info("RunService", "name", req.Msg.ServiceName, "image", req.Msg.Image)
	endpoint := fmt.Sprintf("http://localhost:8080/services/%s", req.Msg.ServiceName)
	return connect.NewResponse(&computev1.RunServiceResponse{EndpointUrl: endpoint}), nil
}

func (s *ComputeServer) TriggerFunction(ctx context.Context, req *connect.Request[computev1.TriggerFunctionRequest]) (*connect.Response[computev1.TriggerFunctionResponse], error) {
	slog.Info("TriggerFunction", "name", req.Msg.FunctionName)
	result := fmt.Sprintf("Function %s executed successfully.", req.Msg.FunctionName)
	return connect.NewResponse(&computev1.TriggerFunctionResponse{Result: result}), nil
}

func main() {
	server := &ComputeServer{}
	mux := http.NewServeMux()
	path, handler := computev1connect.NewComputeServiceHandler(server)
	mux.Handle(path, handler)

	port := "8095" // From genesis.json
	slog.Info("ComputeManager starting", "port", port)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Server failed", "error", err)
	}
}
