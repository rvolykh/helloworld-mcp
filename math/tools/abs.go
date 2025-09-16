package tools

import (
	"context"
	"math"
)

const AbsDescription = `Abs returns the absolute value of x.
Special cases are:

Abs(±Inf) = +Inf
Abs(NaN) = NaN`

type AbsInput struct {
	X float64 `json:"x"`
}

type AbsOutput struct {
	Result float64 `json:"result"`
}

func Abs(ctx context.Context, input AbsInput) (AbsOutput, error) {
	result := math.Abs(input.X)
	return AbsOutput{Result: result}, nil
}
