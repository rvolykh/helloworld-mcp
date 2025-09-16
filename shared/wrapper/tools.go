package wrapper

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ToolFunc[I any, O any] func(context.Context, I) (O, error)

func WrapMCPTool[I any, O any](fn ToolFunc[I, O]) mcp.ToolHandlerFor[I, O] {
	return func(ctx context.Context, _ *mcp.CallToolRequest, i I) (*mcp.CallToolResult, O, error) {
		o, err := fn(ctx, i)
		return nil, o, err
	}
}
