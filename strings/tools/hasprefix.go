package tools

import (
	"context"
	"strings"
)

const HasPrefixDescription = `HasPrefix tests whether the string s begins with prefix.`

type HasPrefixInput struct {
	S      string `json:"s"`
	Prefix string `json:"prefix"`
}

type HasPrefixOutput struct {
	Result bool `json:"result"`
}

func HasPrefix(ctx context.Context, input HasPrefixInput) (HasPrefixOutput, error) {
	result := strings.HasPrefix(input.S, input.Prefix)
	return HasPrefixOutput{Result: result}, nil
}
