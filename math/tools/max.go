package tools

import (
	"context"
	"math"
)

const MaxDescription = `Max returns the larger of x or y.
Special cases are:

Max(x, +Inf) = Max(+Inf, x) = +Inf
Max(x, NaN) = Max(NaN, x) = NaN
Max(+0, ±0) = Max(±0, +0) = +0
Max(-0, -0) = -0`

type MaxInput struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type MaxOutput struct {
	Result float64 `json:"result"`
}

func Max(ctx context.Context, input MaxInput) (MaxOutput, error) {
	result := math.Max(input.X, input.Y)
	return MaxOutput{Result: result}, nil
}
