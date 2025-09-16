package tools

import (
	"context"
	"math"
)

const TruncDescription = `Trunc returns the integer value of x.
Special cases are:

Trunc(±0) = ±0
Trunc(±Inf) = ±Inf
Trunc(NaN) = NaN`

type TruncInput struct {
	X float64 `json:"x"`
}

type TruncOutput struct {
	Result float64 `json:"result"`
}

func Trunc(ctx context.Context, input TruncInput) (TruncOutput, error) {
	result := math.Trunc(input.X)
	return TruncOutput{Result: result}, nil
}
