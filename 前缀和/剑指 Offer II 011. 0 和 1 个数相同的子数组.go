package 前缀和

func findMaxLength(nums []int) int {
	pre := make([]int, len(nums)+1)
	m := map[int]int{}
	m[0] = 0
	ret := 0
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] == 0 {
			pre[i] = pre[i-1] - 1
		} else {
			pre[i] = pre[i-1] + nums[i-1]
		}

		if idx, ok := m[pre[i]]; ok {
			if i-idx > ret {
				ret = i - idx
			}
		} else {
			m[pre[i]] = i
		}
	}

	return ret
}
