package tools

import (
	"context"
	"strings"
)

const RepeatDescription = `Repeat returns a new string consisting of count copies of the string s.
It panics if count is negative or if the result of (len(s) * count) overflows.`

type RepeatInput struct {
	S     string `json:"s"`
	Count int    `json:"count"`
}

type RepeatOutput struct {
	Result string `json:"result"`
}

func Repeat(ctx context.Context, input RepeatInput) (RepeatOutput, error) {
	result := strings.Repeat(input.S, input.Count)
	return RepeatOutput{Result: result}, nil
}
