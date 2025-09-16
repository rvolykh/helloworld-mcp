package tools

import (
	"context"
	"strings"
)

const SplitDescription = `Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
If sep is empty, Split splits after each UTF-8 sequence.
If both s and sep are empty, Split returns an empty slice.`

type SplitInput struct {
	S   string `json:"s"`
	Sep string `json:"sep"`
}

type SplitOutput struct {
	Result []string `json:"result"`
}

func Split(ctx context.Context, input SplitInput) (SplitOutput, error) {
	result := strings.Split(input.S, input.Sep)
	return SplitOutput{Result: result}, nil
}
