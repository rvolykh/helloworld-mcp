package tools

import (
	"context"
	"math"
)

const SinDescription = `Sin returns the sine of the radian argument x.
Special cases are:

Sin(±0) = ±0
Sin(±Inf) = NaN
Sin(NaN) = NaN`

type SinInput struct {
	X float64 `json:"x"`
}

type SinOutput struct {
	Result float64 `json:"result"`
}

func Sin(ctx context.Context, input SinInput) (SinOutput, error) {
	result := math.Sin(input.X)
	return SinOutput{Result: result}, nil
}
