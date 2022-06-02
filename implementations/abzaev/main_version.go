package abzaev

import "math"

func AbzaevSuperFuncImplV1(x1 float64, x2 float64, n uint8) float64 {
	if n == 0 {
		return x1
	}
	if n == 1 {
		return x1 * x2
	}

	var two_multiplication = x1 * x2

	var n1 float64 = 1
	var n2 float64 = 1

	for i := uint8(2); i < n; i++ {
		var temp1 = n1

		n1 = n2
		n2 = temp1 + n2
	}

	var r1 = math.Pow(x1, n1)
	var r2 = math.Pow(two_multiplication, n2)

	return r1 * r2
}
