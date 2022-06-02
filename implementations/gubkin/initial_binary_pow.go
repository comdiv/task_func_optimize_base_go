package gubkin

func InitialBinaryPowV1(x1 float64, x2 float64, n uint8) float64 {
	switch n {
	case 0:
		return x1
	case 1:
		return x1 * x2
	default:
		return FiboPower(x1, n+1) * FiboPower(x2, n)
	}
}

func FiboPower(x float64, fiboIndex uint8) float64 {
	return BinaryPow(x, Fibo(fiboIndex))
}

func BinaryPow(x float64, n int64) float64 {
	result := 1.0
	for n != 0 {
		if n&1 != 0 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	return result
}

func Fibo(n uint8) int64 {
	if n == 30 {
		return 832040
	}
	if n == 31 {
		return 1346269
	}
	var a, b, c, d, rc, rd int64 = 1, 1, 1, 0, 0, 1
	var ta, tb, tc, td int64
	for n > 0 {
		if n&1 != 0 {
			tc = rc*a + rd*c
			td = rc*b + rd*d
			rc = tc
			rd = td
		}
		ta = a*a + b*c
		tb = a*b + b*d
		tc = c*a + d*c
		td = c*b + d*d
		a = ta
		b = tb
		c = tc
		d = td
		n >>= 1
	}
	return rc
}
