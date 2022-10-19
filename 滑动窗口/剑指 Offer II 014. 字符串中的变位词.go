package 滑动窗口

func checkInclusion(s1 string, s2 string) bool {
	m := map[int]int{}
	for _, ch := range s1 {

		m[int(ch)]++
	}

	for l, r := 0, 0; r < len(s2); r++ {
		m[int(s2[r])]--
		for m[int(s2[r])] < 0 {
			m[int(s2[l])]++
			l++
		}

		if r-l+1 == len(s1) {
			return true
		}
	}

	return false
}
