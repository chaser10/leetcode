package 滑动窗口

func minWindow(s string, t string) string {
	start := 0
	minLength := len(s) + 1
	allCh := 0
	findCh := 0
	m1, m2 := map[int]int{}, map[int]int{}
	for _, ch := range t {
		m1[int(ch)]++
		if m1[int(ch)] == 1 {
			allCh++
		}
	}

	for l, r := 0, 0; r < len(s); r++ {
		cur := int(s[r])
		m2[cur]++
		if m2[cur] == m1[cur] {
			findCh++
		}

		for findCh == allCh {
			if r-l+1 < minLength {
				start = l
				minLength = r - l + 1
			}
			front := int(s[l])
			if m2[front] == m1[front] {
				findCh--
			}
			m2[front]--
			l++
		}
	}

	if minLength == len(s)+1 {
		return ""
	}

	return s[start : start+minLength]
}
