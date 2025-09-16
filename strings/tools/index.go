package tools

import (
	"context"
	"strings"
)

const IndexDescription = `Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.`

type IndexInput struct {
	S      string `json:"s"`
	Substr string `json:"substr"`
}

type IndexOutput struct {
	Result int `json:"result"`
}

func Index(ctx context.Context, input IndexInput) (IndexOutput, error) {
	result := strings.Index(input.S, input.Substr)
	return IndexOutput{Result: result}, nil
}
