package tools

import (
	"context"
	"strings"
)

const TrimPrefixDescription = `TrimPrefix returns s without the provided leading prefix string. 
If s doesn't start with prefix, s is returned unchanged.`

type TrimPrefixInput struct {
	S      string `json:"s"`
	Prefix string `json:"prefix"`
}

type TrimPrefixOutput struct {
	Result string `json:"result"`
}

func TrimPrefix(ctx context.Context, input TrimPrefixInput) (TrimPrefixOutput, error) {
	result := strings.TrimPrefix(input.S, input.Prefix)
	return TrimPrefixOutput{Result: result}, nil
}
