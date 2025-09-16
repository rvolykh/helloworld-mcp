package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rvolykh/helloworld-mcp/math/tools"
	"github.com/rvolykh/helloworld-mcp/shared/middleware"
)

var (
	host = flag.String("host", "localhost", "host to connect to/listen on")
	port = flag.Int("port", 8080, "port number to connect to/listen on")
)

func main() {
	out := flag.CommandLine.Output()
	flag.Usage = func() {
		fmt.Fprintf(out, "Usage: %s [-port <port] [-host <host>]\n\n", os.Args[0])
		fmt.Fprintf(out, "This MCP over HTTP for Go math.\n")
		fmt.Fprintf(out, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(out, "\nExamples:\n")
		fmt.Fprintf(out, "  %s -port 9000 -host 0.0.0.0 server\n", os.Args[0])
		os.Exit(1)
	}
	flag.Parse()

	url := fmt.Sprintf("%s:%d", *host, *port)

	// Create a server with tools.
	server := mcp.NewServer(&mcp.Implementation{Name: "math", Version: "v1.0.0"}, nil)

	tools := map[string][]any{
		"Max":   {tools.MaxDescription, tools.Max},
		"Min":   {tools.MinDescription, tools.Min},
		"Abs":   {tools.AbsDescription, tools.Abs},
		"Sqrt":  {tools.SqrtDescription, tools.Sqrt},
		"Pow":   {tools.PowDescription, tools.Pow},
		"Sin":   {tools.SinDescription, tools.Sin},
		"Cos":   {tools.CosDescription, tools.Cos},
		"Tan":   {tools.TanDescription, tools.Tan},
		"Log":   {tools.LogDescription, tools.Log},
		"Log10": {tools.Log10Description, tools.Log10},
		"Exp":   {tools.ExpDescription, tools.Exp},
		"Round": {tools.RoundDescription, tools.Round},
		"Trunc": {tools.TruncDescription, tools.Trunc},
		"Mod":   {tools.ModDescription, tools.Mod},
		"Hypot": {tools.HypotDescription, tools.Hypot},
		"Floor": {tools.FloorDescription, tools.Floor},
		"Ceil":  {tools.CeilDescription, tools.Ceil},
	}
	for name, tool := range tools {
		mcp.AddTool(server, &mcp.Tool{
			Name:        name,
			Description: tool[0].(string),
		}, WrapMathTool(tool[1].(func(context.Context, any) (any, error))))
	}

	// Run the server over HTTP
	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, nil)
	handlerWithLogging := middleware.LoggingHandler(handler)
	log.Printf("MCP server listening on %s", url)

	// Start the HTTP server with logging handler.
	if err := http.ListenAndServe(url, handlerWithLogging); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func WrapMathTool[I any, O any](
	toolFunc func(context.Context, I) (O, error),
) func(context.Context, *mcp.CallToolRequest, I) (*mcp.CallToolResult, O, error) {

	return func(ctx context.Context, _ *mcp.CallToolRequest, i I) (*mcp.CallToolResult, O, error) {
		o, err := toolFunc(ctx, i)
		return nil, o, err
	}
}
