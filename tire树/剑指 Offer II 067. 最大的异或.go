package tire树

type tire map[int]tire

func findMaximumXOR(nums []int) int {
	root := tire{}
	for _, num := range nums {
		cur := root
		for i := 32; i >= 0; i-- {
			if cur[num>>i&1] == nil {
				cur[num>>i&1] = tire{}
			}

			cur = cur[num>>i&1]
		}
	}

	ret := 0
	for _, num := range nums {
		cur := root
		curMax := 0
		for i := 32; i >= 0; i-- {
			high := num >> i & 1
			if cur[1-high] != nil {
				curMax |= 1 << i
				cur = cur[1-high]
			} else {
				cur = cur[high]
			}
		}

		if curMax > ret {
			ret = curMax
		}
	}

	return ret
}
