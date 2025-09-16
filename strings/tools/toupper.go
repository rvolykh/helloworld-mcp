package tools

import (
	"context"
	"strings"
)

const ToUpperDescription = `ToUpper returns s with all Unicode letters mapped to their upper case.`

type ToUpperInput struct {
	S string `json:"s"`
}

type ToUpperOutput struct {
	Result string `json:"result"`
}

func ToUpper(ctx context.Context, input ToUpperInput) (ToUpperOutput, error) {
	result := strings.ToUpper(input.S)
	return ToUpperOutput{Result: result}, nil
}
