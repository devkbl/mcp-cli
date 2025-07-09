package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// input variables
	var port int
	var path, name string

	// read from command line
	flag.IntVar(&port, "port", 8080, "the port number you want your MCP server to run on.")
	flag.StringVar(&path, "path", "mcp", "the endpoint path you want to expose via the server, eg 'mcp'")
	flag.StringVar(&name, "name", "demo-mcp-server", "the name of your MCP server")

	flag.Parse()

	// instantiate new MCP server
	mcpServer := server.NewMCPServer(
		name,
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// attach tools to MCP server
	mcpServer.AddTools(GetTools()...)

	server := server.NewStreamableHTTPServer(mcpServer, server.WithEndpointPath(fmt.Sprintf("/%s", path)))

	slog.Default().Info(fmt.Sprintf("starting %s on localhost:%v", name, port))

	if err := server.Start(fmt.Sprintf(":%v", port)); err != nil {
		slog.Default().Error("error starting MCP HTTP server", slog.Any("error", err))
	}
}

func GetTools() []server.ServerTool {
	tools := []server.ServerTool{
		WeatherTool(),
	}

	return tools
}

func WeatherTool() server.ServerTool {
	tool := mcp.NewTool(
		"weather_tool",
		mcp.WithDescription("returns the weather for a given location."),
		mcp.WithString(
			"location",
			mcp.Required(),
			mcp.Description("Location to return the weather for."),
		),
	)

	handler := func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		location, err := req.RequireString("location")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		// convert to lower to avoid case issues
		location = strings.ToLower(location)

		if location == "united kingdom" || location == "uk" {
			return mcp.NewToolResultText(fmt.Sprintf("The current weather in %s is 26 degrees celsius.", location)), nil
		}

		return mcp.NewToolResultErrorf(fmt.Sprintf("Sorry, we couldn't get the weather for: %s", location)), nil
	}

	return server.ServerTool{
		Tool:    tool,
		Handler: handler,
	}
}
