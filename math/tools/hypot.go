package tools

import (
	"context"
	"math"
)

const HypotDescription = `Hypot returns Sqrt(p*p + q*q), taking care to avoid
unnecessary overflow and underflow.

Special cases are:
Hypot(±Inf, q) = +Inf
Hypot(p, ±Inf) = +Inf
Hypot(NaN, q) = NaN
Hypot(p, NaN) = NaN`

type HypotInput struct {
	P float64 `json:"p"`
	Q float64 `json:"q"`
}

type HypotOutput struct {
	Result float64 `json:"result"`
}

func Hypot(ctx context.Context, input HypotInput) (HypotOutput, error) {
	result := math.Hypot(input.P, input.Q)
	return HypotOutput{Result: result}, nil
}
