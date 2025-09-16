package tools

import (
	"context"
	"math"
)

const CosDescription = `Cos returns the cosine of the radian argument x.
Special cases are:

Cos(±Inf) = NaN
Cos(NaN) = NaN`

type CosInput struct {
	X float64 `json:"x"`
}

type CosOutput struct {
	Result float64 `json:"result"`
}

func Cos(ctx context.Context, input CosInput) (CosOutput, error) {
	result := math.Cos(input.X)
	return CosOutput{Result: result}, nil
}
