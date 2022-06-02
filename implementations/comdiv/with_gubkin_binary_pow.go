package comdiv

import (
	"github.com/comdiv/task_func_optimize_base_go/implementations/commons"
	"github.com/comdiv/task_func_optimize_base_go/implementations/gubkin"
)

func WithGubkinBinaryPowOptimizedV2(x1 float64, x2 float64, n uint8) float64 {
	switch n {
	case 0:
		return x1
	case 1:
		return x1 * x2
	default:
		return gubkin.BinaryPow(x1*x2, commons.FibsI64[n]) * gubkin.BinaryPow(x1, commons.FibsI64[n-1])
	}
}
