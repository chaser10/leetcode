package 二分

func check(piles []int, speed int) int {
	time := 0
	for _, pile := range piles {
		time += (pile + speed - 1) / speed
	}

	return time
}

func minEatingSpeed(piles []int, h int) int {
	left, right := 0, int(1e9)
	for left < right {
		mid := (left + right) >> 1
		if check(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
