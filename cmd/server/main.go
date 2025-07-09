package main

import (
	"flag"
	"fmt"
	"log/slog"

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
	mcpServer := server.NewMCPServer(name, "1.0.0")

	server := server.NewStreamableHTTPServer(mcpServer, server.WithEndpointPath(fmt.Sprintf("/%s", path)))

	slog.Default().Info(fmt.Sprintf("starting %s on localhost:%v", name, port))

	if err := server.Start(fmt.Sprintf(":%v", port)); err != nil {
		slog.Default().Error("error starting MCP HTTP server", slog.Any("error", err))
	}
}
