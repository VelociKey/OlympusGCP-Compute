package main

import (
	"context"
	"olympus.fleet/00SDLC/OlympusForge/70000-Environmental-Harness/dagger/olympusgcp-compute/internal/dagger"
)

type OlympusGCPCompute struct{}

func (m *OlympusGCPCompute) HelloWorld(ctx context.Context) string {
	return "Hello from OlympusGCP-Compute!"
}

func main() {
	dagger.Serve()
}
