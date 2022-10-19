package 滑动窗口

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	ret := len(nums) + 1
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			if r-l+1 < ret {
				ret = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}

	if ret == len(nums)+1 {
		return 0
	}
	return ret
}
