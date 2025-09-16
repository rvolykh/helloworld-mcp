package tools

import (
	"context"
	"strings"
)

const LastIndexDescription = `LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.`

type LastIndexInput struct {
	S      string `json:"s"`
	Substr string `json:"substr"`
}

type LastIndexOutput struct {
	Result int `json:"result"`
}

func LastIndex(ctx context.Context, input LastIndexInput) (LastIndexOutput, error) {
	result := strings.LastIndex(input.S, input.Substr)
	return LastIndexOutput{Result: result}, nil
}
