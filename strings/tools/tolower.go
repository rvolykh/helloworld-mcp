package tools

import (
	"context"
	"strings"
)

const ToLowerDescription = `ToLower returns s with all Unicode letters mapped to their lower case.`

type ToLowerInput struct {
	S string `json:"s"`
}

type ToLowerOutput struct {
	Result string `json:"result"`
}

func ToLower(ctx context.Context, input ToLowerInput) (ToLowerOutput, error) {
	result := strings.ToLower(input.S)
	return ToLowerOutput{Result: result}, nil
}
