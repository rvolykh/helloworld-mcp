package tools

import (
	"context"
	"math"
)

const ModDescription = `Mod returns the floating-point remainder of x/y.
The magnitude of the result is less than y and its
sign agrees with that of x.

Special cases are:
Mod(±Inf, y) = NaN
Mod(NaN, y) = NaN
Mod(x, 0) = NaN
Mod(x, ±Inf) = x
Mod(x, NaN) = NaN`

type ModInput struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ModOutput struct {
	Result float64 `json:"result"`
}

func Mod(ctx context.Context, input ModInput) (ModOutput, error) {
	result := math.Mod(input.X, input.Y)
	return ModOutput{Result: result}, nil
}
