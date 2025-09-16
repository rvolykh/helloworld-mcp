package tools

import (
	"context"
	"math"
)

const SqrtDescription = `Sqrt returns the square root of x.
Special cases are:

Sqrt(+Inf) = +Inf
Sqrt(±0) = ±0
Sqrt(x < 0) = NaN
Sqrt(NaN) = NaN`

type SqrtInput struct {
	X float64 `json:"x"`
}

type SqrtOutput struct {
	Result float64 `json:"result"`
}

func Sqrt(ctx context.Context, input SqrtInput) (SqrtOutput, error) {
	result := math.Sqrt(input.X)
	return SqrtOutput{Result: result}, nil
}
