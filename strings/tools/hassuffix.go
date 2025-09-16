package tools

import (
	"context"
	"strings"
)

const HasSuffixDescription = `HasSuffix tests whether the string s ends with suffix.`

type HasSuffixInput struct {
	S      string `json:"s"`
	Suffix string `json:"suffix"`
}

type HasSuffixOutput struct {
	Result bool `json:"result"`
}

func HasSuffix(ctx context.Context, input HasSuffixInput) (HasSuffixOutput, error) {
	result := strings.HasSuffix(input.S, input.Suffix)
	return HasSuffixOutput{Result: result}, nil
}
