package tools

import (
	"context"
	"strings"
)

const CompareDescription = `Compare returns an integer comparing two strings lexicographically.
The result will be 0 if a == b, -1 if a < b, and +1 if a > b.

Use Compare when you need to perform a three-way comparison (with slices.SortFunc, for example).
It is usually clearer and always faster to use the built-in string comparison operators ==, <, >, and so on.`

type CompareInput struct {
	A string `json:"a"`
	B string `json:"b"`
}

type CompareOutput struct {
	Result int `json:"result"`
}

func Compare(ctx context.Context, input CompareInput) (CompareOutput, error) {
	result := strings.Compare(input.A, input.B)
	return CompareOutput{Result: result}, nil
}
