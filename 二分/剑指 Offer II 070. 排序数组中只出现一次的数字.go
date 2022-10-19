package 二分

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == nums[mid+1] {
			if mid%2 == 0 {
				left = mid + 2
			} else {
				right = mid - 1
			}
		} else if nums[mid] == nums[mid-1] {
			if mid%2 == 1 {
				left = mid + 1
			} else {
				right = mid - 2
			}
		} else {
			return nums[mid]
		}
	}

	return nums[left]
}
