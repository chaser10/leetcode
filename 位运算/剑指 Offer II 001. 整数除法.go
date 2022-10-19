package 位运算

import (
	"math"
)

func divide(a int, b int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}
	flag := false
	ret := 0

	if a > 0 && b < 0 || a < 0 && b > 0 {
		flag = true
	}
	if b == math.MinInt32 {
		if a == math.MinInt32 {
			return 1
		} else {
			return 0
		}
	}

	if a == math.MinInt32 {
		if b == 1 {
			return a
		} else if b == -1 {
			return math.MaxInt32
		} else {
			ret += 1
			a -= -abs(b)
		}
	}

	a = abs(a)
	b = abs(b)

	for i := 32; i >= 0; i-- {
		if a>>i >= b {
			ret += 1 << i
			a -= b << i
		}
	}

	if flag {
		ret = -ret
	}
	return ret
}
