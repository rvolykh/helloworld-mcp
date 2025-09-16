package tools

import (
	"context"
	"strings"
)

const JoinDescription = `Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.`

type JoinInput struct {
	Elems []string `json:"elems"`
	Sep   string   `json:"sep"`
}

type JoinOutput struct {
	Result string `json:"result"`
}

func Join(ctx context.Context, input JoinInput) (JoinOutput, error) {
	result := strings.Join(input.Elems, input.Sep)
	return JoinOutput{Result: result}, nil
}
