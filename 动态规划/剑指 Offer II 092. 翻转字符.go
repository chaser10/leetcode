func minFlipsMonoIncr(s string) int {
	n := len(s)
	f := make([]int, 2)
	for i := 1; i <= n; i++ {
		if s[i - 1] == '0' {
			f[1] = min(f[0], f[1]) + 1
			f[0] = f[0]
		} else {
			f[1] = min(f[0], f[1])
			f[0] = f[0] + 1
		}
	} 

	return min(f[0], f[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}