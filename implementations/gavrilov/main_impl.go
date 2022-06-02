package gavrilov

func GavrilovSuperFuncImplV1(x1 float64, x2 float64, n uint8) float64 {
	switch n {
	case 0:
		return x1
	default:
		var acc float64 = x1
		var res float64 = x1 * x2
		for i := uint8(2); i <= n; i++ {
			acc, res = res, acc*res
		}
		return res
	}
}
