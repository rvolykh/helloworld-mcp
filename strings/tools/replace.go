package tools

import (
	"context"
	"strings"
)

const ReplaceDescription = `Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. 
If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string. 
If n < 0, there is no limit on the number of replacements.`

type ReplaceInput struct {
	S   string `json:"s"`
	Old string `json:"old"`
	New string `json:"new"`
	N   int    `json:"n"`
}

type ReplaceOutput struct {
	Result string `json:"result"`
}

func Replace(ctx context.Context, input ReplaceInput) (ReplaceOutput, error) {
	result := strings.Replace(input.S, input.Old, input.New, input.N)
	return ReplaceOutput{Result: result}, nil
}
