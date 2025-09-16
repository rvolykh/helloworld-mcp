package tools

import (
	"context"
	"strings"
)

const ContainsDescription = `Contains reports whether substr is within s.`

type ContainsInput struct {
	S      string `json:"s"`
	Substr string `json:"substr"`
}

type ContainsOutput struct {
	Result bool `json:"result"`
}

func Contains(ctx context.Context, input ContainsInput) (ContainsOutput, error) {
	result := strings.Contains(input.S, input.Substr)
	return ContainsOutput{Result: result}, nil
}
