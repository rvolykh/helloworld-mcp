package tools

import (
	"context"
	"strings"
)

const EqualFoldDescription = `EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under simple Unicode case-folding, 
which is a more general form of case-insensitivity.`

type EqualFoldInput struct {
	S string `json:"s"`
	T string `json:"t"`
}

type EqualFoldOutput struct {
	Result bool `json:"result"`
}

func EqualFold(ctx context.Context, input EqualFoldInput) (EqualFoldOutput, error) {
	result := strings.EqualFold(input.S, input.T)
	return EqualFoldOutput{Result: result}, nil
}
