package implementations

import (
	"github.com/comdiv/task_func_optimize_base_go/basis"
	"github.com/comdiv/task_func_optimize_base_go/implementations/abzaev"
	"github.com/comdiv/task_func_optimize_base_go/implementations/comdiv"
	"github.com/comdiv/task_func_optimize_base_go/implementations/gavrilov"
	"github.com/comdiv/task_func_optimize_base_go/implementations/gubkin"
	"github.com/comdiv/task_func_optimize_base_go/implementations/harisov"
	"github.com/comdiv/task_func_optimize_base_go/implementations/pirozhkov"
	"github.com/comdiv/task_func_optimize_base_go/implementations/tetyuev"
)

var AllImplementations = []*basis.FuncDesc{
	{"comdiv    v1 (fib[]+math.Pow)", comdiv.InitialImpl},
	{"gubkin    v1 (fib+BinaryPow)", gubkin.InitialBinaryPowV1},
	{"gubkin    v2 (fib[]+BinaryPow)", gubkin.OptimizedWithWithFibArrayV2},
	{"comdiv    v2 (fib[]+BinaryPow(optimal))", comdiv.WithGubkinBinaryPowOptimizedV2},
	{"tetyuev   v1 (linear)", tetyuev.TetyuevSuperFuncImplV1},
	{"pirozhkov v1 (fib+math.Pow)", pirozhkov.PirozhkovSuperFuncImplV1},
	{"tetyuev   v2 (fib[]+math.Pow)", tetyuev.TetyuevSuperFuncImplV2},
	{"harisov   v1 (fib+SqrtPower)", harisov.MySuperFuncImpl}, // тесты на погрешность не проходят!!!
	{"gavrilov  v1 (linear)", gavrilov.GavrilovSuperFuncImplV1},
	{"abzaev    v1 (fib+math.Pow)", abzaev.AbzaevSuperFuncImplV1},
	{"gubkin    v3 (fib[]+CachedPow)", gubkin.GubkinSuperFuncImplV3},
	{"gubkin    v4 (CachedPow)", gubkin.Gubkin3SuperFuncImplV4},
	{"comdiv    v3 (fib[]+CachedPow+BinaryPowDual)", comdiv.BestOfAllSuperFuncImpl},
}
