package 前缀和

func pivotIndex(nums []int) int {
	sum := 0
	pre := 0
	for _, num := range nums {
		sum += num
	}

	for i, num := range nums {
		if pre == sum-pre-num {
			return i
		}

		pre += num
	}

	return -1
}
