func rob(nums []int) int {
	n := len(nums) - 1
	f1 := make([]int, n)
	f2 := make([]int, n)
	nums1 := nums[1:]
	nums2 := nums[:len(nums) - 1]
	f1[1] = nums[0]
	f2[1] = nums[1]
	for i := 2; i < n; i++ {
		f1[i] = max(f1[i - 1], f1[i - 2] + nums1[i - 1])
		f2[i] = max(f2[i - 1], f2[i - 2] + nums2[i - 1])
	}

	return max(f1[n - 1], f2[n - 1])
}

func max(a, b int) int {
	if a > b { 
		return a
	}

	return b
}