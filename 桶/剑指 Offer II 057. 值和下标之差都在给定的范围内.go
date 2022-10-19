package 桶

import "fmt"

func getId(x int, w int) int {
	/*
		比如窗口是4  那么 1234放在一个桶（id为0）
		-1 -2 -3 -4 放在id为-1的桶里
	*/
	if x > 0 {
		return x / w
	}

	return (x+1)/w - 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	/*
		每次检查当前桶和相邻桶
	*/
	m := map[int]int{}
	for left, right := 0, 0; right < len(nums); right++ {
		if right-left+1 > k+1 {
			delete(m, getId(nums[left], k+1))
		}

		id := getId(nums[right], k+1)
		if _, ok := m[id-1]; ok {
			fmt.Println(nums[right])
			return true
		}

		if v, ok := m[id-1]; ok && abs(v-nums[right]) <= t {
			fmt.Println(nums[right])
			return true
		}

		if v, ok := m[id+1]; ok && abs(v-nums[right]) <= t {
			fmt.Println(nums[right])
			return true
		}

		m[id] = nums[right]
	}

	return false
}
