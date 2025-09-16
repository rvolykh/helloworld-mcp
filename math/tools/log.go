package tools

import (
	"context"
	"math"
)

const LogDescription = `Log returns the natural logarithm of x.
Special cases are:

Log(+Inf) = +Inf
Log(0) = -Inf
Log(x < 0) = NaN
Log(NaN) = NaN`

type LogInput struct {
	X float64 `json:"x"`
}

type LogOutput struct {
	Result float64 `json:"result"`
}

func Log(ctx context.Context, input LogInput) (LogOutput, error) {
	result := math.Log(input.X)
	return LogOutput{Result: result}, nil
}
