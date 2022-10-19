package 思维

func minimizeArrayValue(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	for nums[0] <= max {

	}
}
