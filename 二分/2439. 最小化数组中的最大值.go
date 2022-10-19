package 二分

func check(nums []int, k int) bool {
	have := 0
	for _, num := range nums {
		if num <= k {
			have += k - num
		} else {
			if num-k > have {
				return false
			} else {
				have -= num - k
			}
		}
	}

	return true
}

func minimizeArrayValue(nums []int) int {
	left, right := 0, int(1e9)
	for left < right {
		mid := (left + right) >> 1
		if check(nums, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
