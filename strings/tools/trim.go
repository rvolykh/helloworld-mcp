package tools

import (
	"context"
	"strings"
)

const TrimDescription = `Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.`

type TrimInput struct {
	S      string `json:"s"`
	Cutset string `json:"cutset"`
}

type TrimOutput struct {
	Result string `json:"result"`
}

func Trim(ctx context.Context, input TrimInput) (TrimOutput, error) {
	result := strings.Trim(input.S, input.Cutset)
	return TrimOutput{Result: result}, nil
}
