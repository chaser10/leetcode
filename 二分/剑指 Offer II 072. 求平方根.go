package 二分

func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := (left + right + 1) >> 1
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}
