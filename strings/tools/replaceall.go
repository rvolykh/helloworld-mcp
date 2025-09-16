package tools

import (
	"context"
	"strings"
)

const ReplaceAllDescription = `ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new. 
If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.`

type ReplaceAllInput struct {
	S   string `json:"s"`
	Old string `json:"old"`
	New string `json:"new"`
}

type ReplaceAllOutput struct {
	Result string `json:"result"`
}

func ReplaceAll(ctx context.Context, input ReplaceAllInput) (ReplaceAllOutput, error) {
	result := strings.ReplaceAll(input.S, input.Old, input.New)
	return ReplaceAllOutput{Result: result}, nil
}
