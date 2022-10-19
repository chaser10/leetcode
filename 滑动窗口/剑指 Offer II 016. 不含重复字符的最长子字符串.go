package 滑动窗口

func lengthOfLongestSubstring(s string) int {
	m := map[int]int{}
	ret := 0
	for l, r := 0, 0; r < len(s); r++ {
		m[int(s[r])]++
		for m[int(s[r])] > 1 {
			m[int(s[l])]--
			l++
		}

		if r-l+1 > ret {
			ret = r - l + 1
		}
	}

	return ret
}
