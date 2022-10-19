package 双指针

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left, right}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
