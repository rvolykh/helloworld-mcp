package tools

import (
	"context"
	"math"
)

const MinDescription = `Min returns the smaller of x or y.
Special cases are:

Min(x, -Inf) = Min(-Inf, x) = -Inf
Min(x, NaN) = Min(NaN, x) = NaN
Min(-0, ±0) = Min(±0, -0) = -0`

type MinInput struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type MinOutput struct {
	Result float64 `json:"result"`
}

func Min(ctx context.Context, input MinInput) (MinOutput, error) {
	result := math.Min(input.X, input.Y)
	return MinOutput{Result: result}, nil
}
