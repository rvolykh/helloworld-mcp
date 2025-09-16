package tools

import (
	"context"
	"strings"
)

const TrimSuffixDescription = `TrimSuffix returns s without the provided trailing suffix string. 
If s doesn't end with suffix, s is returned unchanged.`

type TrimSuffixInput struct {
	S      string `json:"s"`
	Suffix string `json:"suffix"`
}

type TrimSuffixOutput struct {
	Result string `json:"result"`
}

func TrimSuffix(ctx context.Context, input TrimSuffixInput) (TrimSuffixOutput, error) {
	result := strings.TrimSuffix(input.S, input.Suffix)
	return TrimSuffixOutput{Result: result}, nil
}
