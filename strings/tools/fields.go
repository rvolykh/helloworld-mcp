package tools

import (
	"context"
	"strings"
)

const FieldsDescription = `Fields splits the string s around each instance of one or more consecutive white space characters, 
as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.`

type FieldsInput struct {
	S string `json:"s"`
}

type FieldsOutput struct {
	Result []string `json:"result"`
}

func Fields(ctx context.Context, input FieldsInput) (FieldsOutput, error) {
	result := strings.Fields(input.S)
	return FieldsOutput{Result: result}, nil
}
