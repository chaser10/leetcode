package 滑动窗口

func numSubarrayProductLessThanK(nums []int, k int) int {
	mutilply := 1
	ret := 0
	for l, r := 0, 0; r < len(nums); r++ {
		mutilply *= nums[r]
		for l <= r && mutilply >= k {
			mutilply /= nums[l]
			l++
		}
		if mutilply < k {
			ret += r - l + 1
		}
	}

	return ret
}
