package 二分

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] < target {
		return len(nums) - 1
	}
	return left
}
