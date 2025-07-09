package mcp

import (
	"context"
	"log/slog"

	"github.com/mark3labs/mcp-go/client"
)

func NewMCPClient(ctx context.Context, url string) *client.Client {
	slog.Debug("creating client", slog.String("url", url))
	// basic http streaming client
	// TODO: get from flags
	return mcpClient
}
