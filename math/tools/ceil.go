package tools

import (
	"context"
	"math"
)

const CeilDescription = `Ceil returns the least integer value greater than or equal to x.
Special cases are:

Ceil(±0) = ±0
Ceil(±Inf) = ±Inf
Ceil(NaN) = NaN`

type CeilInput struct {
	X float64 `json:"x"`
}

type CeilOutput struct {
	Result float64 `json:"result"`
}

func Ceil(ctx context.Context, input CeilInput) (CeilOutput, error) {
	result := math.Ceil(input.X)
	return CeilOutput{Result: result}, nil
}
