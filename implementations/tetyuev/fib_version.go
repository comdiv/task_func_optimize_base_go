package tetyuev

import (
	"github.com/comdiv/task_func_optimize_base_go/implementations/commons"
	"math"
)

func TetyuevSuperFuncImplV2(x1 float64, x2 float64, n uint8) float64 {
	switch n {
	case 0:
		return x1
	case 1:
		return x1 * x2
	}
	return math.Pow(x1, float64(commons.FibsI64[n+1])) * math.Pow(x2, float64(commons.FibsI64[n]))
}
