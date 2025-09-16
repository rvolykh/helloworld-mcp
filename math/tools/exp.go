package tools

import (
	"context"
	"math"
)

const ExpDescription = `Exp returns e**x, the base-e exponential of x.
Special cases are:

Exp(+Inf) = +Inf
Exp(NaN) = NaN

Very large values overflow to 0 or +Inf.
Very small values underflow to 1.`

type ExpInput struct {
	X float64 `json:"x"`
}

type ExpOutput struct {
	Result float64 `json:"result"`
}

func Exp(ctx context.Context, input ExpInput) (ExpOutput, error) {
	result := math.Exp(input.X)
	return ExpOutput{Result: result}, nil
}
