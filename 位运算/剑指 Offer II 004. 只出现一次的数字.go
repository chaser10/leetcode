package 位运算

func singleNumber(nums []int) int {
	ret := int32(0)
	for i := 0; i < 32; i++ {
		cur := 0
		for _, num := range nums {
			cur += (num >> i) & 1
		}

		ret += (cur % 3) << i
	}

	return int(ret)
}
