package 字符串

func countSubstrings(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += expand(i, i, s)
		ans += expand(i, i+1, s)
	}

	return ans
}

func expand(i int, j int, s string) int {
	ret := 0
	for i >= 0 && j < len(s) {
		if s[i] == s[j] {
			ret++
		} else {
			break
		}

		i--
		j++
	}

	return ret
}
