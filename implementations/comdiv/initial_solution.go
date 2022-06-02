package comdiv

import (
	"github.com/comdiv/task_func_optimize_base_go/implementations/commons"
	"math"
)

func InitialImpl(x1 float64, x2 float64, n uint8) float64 {
	switch n {
	case 0:
		return x1
	case 1:
		return x1 * x2
	default:
		return math.Pow(x1*x2, commons.FibsF64[n]) * math.Pow(x1, commons.FibsF64[n-1])
	}
}
