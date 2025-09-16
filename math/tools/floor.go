package tools

import (
	"context"
	"math"
)

const FloorDescription = `Floor returns the greatest integer value less than or equal to x.
Special cases are:

Floor(±0) = ±0
Floor(±Inf) = ±Inf
Floor(NaN) = NaN`

type FloorInput struct {
	X float64 `json:"x"`
}

type FloorOutput struct {
	Result float64 `json:"result"`
}

func Floor(ctx context.Context, input FloorInput) (FloorOutput, error) {
	result := math.Floor(input.X)
	return FloorOutput{Result: result}, nil
}
