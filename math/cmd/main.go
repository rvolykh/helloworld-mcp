package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rvolykh/helloworld-mcp/math/tools"
	"github.com/rvolykh/helloworld-mcp/shared/middleware"
	"github.com/rvolykh/helloworld-mcp/shared/wrapper"
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

	// Create a server
	server := mcp.NewServer(&mcp.Implementation{Name: "math", Version: "v1.0.0"}, nil)

	// Add tools
	mcp.AddTool(server, &mcp.Tool{Name: "Abs", Description: tools.AbsDescription}, wrapper.WrapMCPTool(tools.Abs))
	mcp.AddTool(server, &mcp.Tool{Name: "Ceil", Description: tools.CeilDescription}, wrapper.WrapMCPTool(tools.Ceil))
	mcp.AddTool(server, &mcp.Tool{Name: "Cos", Description: tools.CosDescription}, wrapper.WrapMCPTool(tools.Cos))
	mcp.AddTool(server, &mcp.Tool{Name: "Exp", Description: tools.ExpDescription}, wrapper.WrapMCPTool(tools.Exp))
	mcp.AddTool(server, &mcp.Tool{Name: "Floor", Description: tools.FloorDescription}, wrapper.WrapMCPTool(tools.Floor))
	mcp.AddTool(server, &mcp.Tool{Name: "Hypot", Description: tools.HypotDescription}, wrapper.WrapMCPTool(tools.Hypot))
	mcp.AddTool(server, &mcp.Tool{Name: "Log", Description: tools.LogDescription}, wrapper.WrapMCPTool(tools.Log))
	mcp.AddTool(server, &mcp.Tool{Name: "Log10", Description: tools.Log10Description}, wrapper.WrapMCPTool(tools.Log10))
	mcp.AddTool(server, &mcp.Tool{Name: "Max", Description: tools.MaxDescription}, wrapper.WrapMCPTool(tools.Max))
	mcp.AddTool(server, &mcp.Tool{Name: "Min", Description: tools.MinDescription}, wrapper.WrapMCPTool(tools.Min))
	mcp.AddTool(server, &mcp.Tool{Name: "Mod", Description: tools.ModDescription}, wrapper.WrapMCPTool(tools.Mod))
	mcp.AddTool(server, &mcp.Tool{Name: "Pow", Description: tools.PowDescription}, wrapper.WrapMCPTool(tools.Pow))
	mcp.AddTool(server, &mcp.Tool{Name: "Round", Description: tools.RoundDescription}, wrapper.WrapMCPTool(tools.Round))
	mcp.AddTool(server, &mcp.Tool{Name: "Sin", Description: tools.SinDescription}, wrapper.WrapMCPTool(tools.Sin))
	mcp.AddTool(server, &mcp.Tool{Name: "Sqrt", Description: tools.SqrtDescription}, wrapper.WrapMCPTool(tools.Sqrt))
	mcp.AddTool(server, &mcp.Tool{Name: "Tan", Description: tools.TanDescription}, wrapper.WrapMCPTool(tools.Tan))
	mcp.AddTool(server, &mcp.Tool{Name: "Trunc", Description: tools.TruncDescription}, wrapper.WrapMCPTool(tools.Trunc))

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
