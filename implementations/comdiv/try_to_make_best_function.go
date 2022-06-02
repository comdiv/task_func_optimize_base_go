package comdiv

import (
	"github.com/comdiv/task_func_optimize_base_go/implementations/commons"
	"github.com/comdiv/task_func_optimize_base_go/implementations/gubkin"
)

func BestOfAllSuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	if n >= 31 {
		return SpecialDualVariableFibbonachiBinaryPower(x1, x2, n)
	}
	if n >= 16 {
		return gubkin.FiboPowers[n+1](x1) * gubkin.FiboPowers[n](x2)
	}
	if n == 0 {
		return x1
	}
	x2 = x1 * x2
	for i := uint8(2); i <= n; i++ {
		x1, x2 = x2, x1*x2
	}
	return x2
}

// быстрое возведение в целую положительную степень
func SpecialDualVariableFibbonachiBinaryPower(x1 float64, x2 float64, n uint8) float64 {
	result := 1.0
	nx1 := commons.FibsI64[n+1]
	nx2 := commons.FibsI64[n]
	for nx1 != 0 {
		if nx2&1 != 0 {
			result *= x2
		}
		if nx1&1 != 0 {
			result *= x1
		}
		x1 *= x1
		x2 *= x2
		nx1 >>= 1
		nx2 >>= 1
	}
	return result
}
