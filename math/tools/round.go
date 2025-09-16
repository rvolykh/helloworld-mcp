package tools

import (
	"context"
	"math"
)

const RoundDescription = `Round returns the nearest integer, rounding half values away from zero.
Special cases are:

Round(±0) = ±0
Round(±Inf) = ±Inf
Round(NaN) = NaN`

type RoundInput struct {
	X float64 `json:"x"`
}

type RoundOutput struct {
	Result float64 `json:"result"`
}

func Round(ctx context.Context, input RoundInput) (RoundOutput, error) {
	result := math.Round(input.X)
	return RoundOutput{Result: result}, nil
}
