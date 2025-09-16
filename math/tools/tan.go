package tools

import (
	"context"
	"math"
)

const TanDescription = `Tan returns the tangent of the radian argument x.
Special cases are:

Tan(±0) = ±0
Tan(±Inf) = NaN
Tan(NaN) = NaN`

type TanInput struct {
	X float64 `json:"x"`
}

type TanOutput struct {
	Result float64 `json:"result"`
}

func Tan(ctx context.Context, input TanInput) (TanOutput, error) {
	result := math.Tan(input.X)
	return TanOutput{Result: result}, nil
}
