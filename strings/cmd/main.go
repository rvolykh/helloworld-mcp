package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rvolykh/helloworld-mcp/shared/middleware"
	"github.com/rvolykh/helloworld-mcp/shared/wrapper"
	"github.com/rvolykh/helloworld-mcp/strings/tools"
)

var (
	host = flag.String("host", "localhost", "host to connect to/listen on")
	port = flag.Int("port", 8080, "port number to connect to/listen on")
)

func main() {
	out := flag.CommandLine.Output()
	flag.Usage = func() {
		fmt.Fprintf(out, "Usage: %s [-port <port] [-host <host>]\n\n", os.Args[0])
		fmt.Fprintf(out, "This MCP server provides Go strings functions over HTTP.\n")
		fmt.Fprintf(out, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(out, "\nExamples:\n")
		fmt.Fprintf(out, "  %s -port 9000 -host 0.0.0.0 server\n", os.Args[0])
		os.Exit(1)
	}
	flag.Parse()

	url := fmt.Sprintf("%s:%d", *host, *port)

	// Create a server
	server := mcp.NewServer(&mcp.Implementation{Name: "strings", Version: "v1.0.0"}, nil)

	// Add tools
	mcp.AddTool(server, &mcp.Tool{Name: "Compare", Description: tools.CompareDescription}, wrapper.WrapMathTool(tools.Compare))
	mcp.AddTool(server, &mcp.Tool{Name: "Contains", Description: tools.ContainsDescription}, wrapper.WrapMathTool(tools.Contains))
	mcp.AddTool(server, &mcp.Tool{Name: "Count", Description: tools.CountDescription}, wrapper.WrapMathTool(tools.Count))
	mcp.AddTool(server, &mcp.Tool{Name: "EqualFold", Description: tools.EqualFoldDescription}, wrapper.WrapMathTool(tools.EqualFold))
	mcp.AddTool(server, &mcp.Tool{Name: "Fields", Description: tools.FieldsDescription}, wrapper.WrapMathTool(tools.Fields))
	mcp.AddTool(server, &mcp.Tool{Name: "HasPrefix", Description: tools.HasPrefixDescription}, wrapper.WrapMathTool(tools.HasPrefix))
	mcp.AddTool(server, &mcp.Tool{Name: "HasSuffix", Description: tools.HasSuffixDescription}, wrapper.WrapMathTool(tools.HasSuffix))
	mcp.AddTool(server, &mcp.Tool{Name: "Index", Description: tools.IndexDescription}, wrapper.WrapMathTool(tools.Index))
	mcp.AddTool(server, &mcp.Tool{Name: "Join", Description: tools.JoinDescription}, wrapper.WrapMathTool(tools.Join))
	mcp.AddTool(server, &mcp.Tool{Name: "LastIndex", Description: tools.LastIndexDescription}, wrapper.WrapMathTool(tools.LastIndex))
	mcp.AddTool(server, &mcp.Tool{Name: "Repeat", Description: tools.RepeatDescription}, wrapper.WrapMathTool(tools.Repeat))
	mcp.AddTool(server, &mcp.Tool{Name: "Replace", Description: tools.ReplaceDescription}, wrapper.WrapMathTool(tools.Replace))
	mcp.AddTool(server, &mcp.Tool{Name: "ReplaceAll", Description: tools.ReplaceAllDescription}, wrapper.WrapMathTool(tools.ReplaceAll))
	mcp.AddTool(server, &mcp.Tool{Name: "Split", Description: tools.SplitDescription}, wrapper.WrapMathTool(tools.Split))
	mcp.AddTool(server, &mcp.Tool{Name: "ToLower", Description: tools.ToLowerDescription}, wrapper.WrapMathTool(tools.ToLower))
	mcp.AddTool(server, &mcp.Tool{Name: "ToUpper", Description: tools.ToUpperDescription}, wrapper.WrapMathTool(tools.ToUpper))
	mcp.AddTool(server, &mcp.Tool{Name: "Trim", Description: tools.TrimDescription}, wrapper.WrapMathTool(tools.Trim))
	mcp.AddTool(server, &mcp.Tool{Name: "TrimPrefix", Description: tools.TrimPrefixDescription}, wrapper.WrapMathTool(tools.TrimPrefix))
	mcp.AddTool(server, &mcp.Tool{Name: "TrimSpace", Description: tools.TrimSpaceDescription}, wrapper.WrapMathTool(tools.TrimSpace))
	mcp.AddTool(server, &mcp.Tool{Name: "TrimSuffix", Description: tools.TrimSuffixDescription}, wrapper.WrapMathTool(tools.TrimSuffix))

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
