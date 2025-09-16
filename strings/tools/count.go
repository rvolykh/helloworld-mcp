package tools

import (
	"context"
	"strings"
)

const CountDescription = `Count counts the number of non-overlapping instances of substr in s. 
If substr is an empty string, Count returns 1 + the number of Unicode code points in s.`

type CountInput struct {
	S      string `json:"s"`
	Substr string `json:"substr"`
}

type CountOutput struct {
	Result int `json:"result"`
}

func Count(ctx context.Context, input CountInput) (CountOutput, error) {
	result := strings.Count(input.S, input.Substr)
	return CountOutput{Result: result}, nil
}
