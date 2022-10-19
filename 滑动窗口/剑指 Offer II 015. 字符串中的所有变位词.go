package 滑动窗口

func findAnagrams(s string, p string) []int {
	ret := []int{}
	m := map[int]int{}
	for _, ch := range p {
		m[int(ch)]++
	}

	for l, r := 0, 0; r < len(s); r++ {
		m[int(s[r])]--
		for m[int(s[r])] < 0 {
			m[int(s[l])]++
			l++
		}

		if r-l+1 == len(p) {
			ret = append(ret, l)
		}
	}

	return ret
}
