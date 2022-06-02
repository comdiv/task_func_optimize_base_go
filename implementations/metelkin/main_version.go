package metelkin

func MetelkinSuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	if n == 0 {
		return x1
	} else if n == 1 {
		return x1 * x2
	} else {
		var preLastValue = x1
		var lastValue = x1 * x2

		var currentValue float64
		for i := uint8(2); i <= n; i++ {
			currentValue = preLastValue * lastValue

			preLastValue = lastValue
			lastValue = currentValue
		}

		return currentValue
	}
}
