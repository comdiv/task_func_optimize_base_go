package gubkin

import "github.com/comdiv/task_func_optimize_base_go/implementations/commons"

func OptimizedWithWithFibArrayV2(x1 float64, x2 float64, n uint8) float64 {
	switch n {
	case 0:
		return x1
	case 1:
		return x1 * x2
	default:
		return BinaryPow(x1, commons.FibsI64[n+1]) * BinaryPow(x2, commons.FibsI64[n])
	}
}
