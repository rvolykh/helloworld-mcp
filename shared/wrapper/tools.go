package wrapper

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func WrapMathTool[I any, O any](
	toolFunc func(context.Context, I) (O, error),
) func(context.Context, *mcp.CallToolRequest, I) (*mcp.CallToolResult, O, error) {

	return func(ctx context.Context, _ *mcp.CallToolRequest, i I) (*mcp.CallToolResult, O, error) {
		o, err := toolFunc(ctx, i)
		return nil, o, err
	}
}
