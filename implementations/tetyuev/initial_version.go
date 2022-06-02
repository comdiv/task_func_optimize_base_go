package tetyuev

func TetyuevSuperFuncImplV1(x1 float64, x2 float64, n uint8) float64 {
	// Dynamic programming, bottom-up with heap optimization.

	memF := x1
	memS := x1 * x2

	switch n {
	case 0:
		return memF
	case 1:
		return memS
	}

	answ := float64(0)

	for i := uint8(2); i <= n; i++ {
		answ = memF * memS
		memF, memS = memS, answ
	}

	return answ
}
