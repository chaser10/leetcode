package 前缀和

func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	pre := 0
	m[0] = 1
	ret := 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if cnt, ok := m[pre-k]; ok {
			ret += cnt
		}
		m[pre]++
	}

	return ret
}
