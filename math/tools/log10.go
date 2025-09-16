package tools

import (
	"context"
	"math"
)

const Log10Description = `Log10 returns the decimal logarithm of x.
The special cases are the same as for Log.

Special cases are:
Log10(+Inf) = +Inf
Log10(0) = -Inf
Log10(x < 0) = NaN
Log10(NaN) = NaN`

type Log10Input struct {
	X float64 `json:"x"`
}

type Log10Output struct {
	Result float64 `json:"result"`
}

func Log10(ctx context.Context, input Log10Input) (Log10Output, error) {
	result := math.Log10(input.X)
	return Log10Output{Result: result}, nil
}
