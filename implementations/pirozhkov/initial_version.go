package pirozhkov

import "math"

func PirozhkovSuperFuncImplV1(x1 float64, x2 float64, n uint8) float64 {
	var x1_pow float64 = 1
	var x2_pow float64 = 0
	var i uint8 = 1
	for ; i <= n; i++ {
		tmp := x2_pow
		x2_pow = x1_pow
		x1_pow += tmp
	}
	return math.Pow(x1, x1_pow) * math.Pow(x2, x2_pow)
}
