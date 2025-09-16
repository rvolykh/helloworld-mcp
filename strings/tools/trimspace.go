package tools

import (
	"context"
	"strings"
)

const TrimSpaceDescription = `TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.`

type TrimSpaceInput struct {
	S string `json:"s"`
}

type TrimSpaceOutput struct {
	Result string `json:"result"`
}

func TrimSpace(ctx context.Context, input TrimSpaceInput) (TrimSpaceOutput, error) {
	result := strings.TrimSpace(input.S)
	return TrimSpaceOutput{Result: result}, nil
}
