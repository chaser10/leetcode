package 二分

func guess(num int) int {
	return 0
}
func guessNumber(n int) int {
	left, right := 1, n
	for left < right {
		mid := (left + right) >> 1
		g := guess(mid)
		if g == 0 {
			return mid
		} else if g > 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
