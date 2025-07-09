package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	ctx := context.Background()
	// basic http streaming client
	// TODO: get from flags
	url := "http://localhost:8080/mcp"
	transport, err := transport.NewStreamableHTTP(url)
	if err != nil {
		slog.Default().ErrorContext(ctx, "error creating StreamableHTTP transport", slog.Any("error", err))
	}

	// create new client based on transport
	mcpClient := client.NewClient(transport)

	// close client connection when function ends
	defer func() {
		if err := mcpClient.Close(); err != nil {
			slog.Default().ErrorContext(ctx, "error closing client", slog.Any("error", err))
		}
	}()

	// initialize client

	initReq := mcp.InitializeRequest{}

	_, err = mcpClient.Initialize(ctx, initReq)
	if err != nil {
		slog.Default().ErrorContext(ctx, "error initializing client session", slog.Any("error", err))
	}

	// list tools
	toolReq := mcp.ListToolsRequest{}
	tools, err := mcpClient.ListTools(ctx, toolReq)
	if err != nil {
		slog.Default().ErrorContext(ctx, "error listing MCP tools", slog.Any("error", err))
	}

	for _, tool := range tools.Tools {
		fmt.Println(tool)
	}
}
