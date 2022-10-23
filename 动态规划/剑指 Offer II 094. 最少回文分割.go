func minCut(s string) int {
	n := len(s)
    f := make([][]bool, n)
    for i := range f {
        f[i] = make([]bool, n)
        for j := range f[i] {
            if i == j {
                f[i][j] = true
            }
        }
    }

	for i := 2; i <= n; i++ {
		for j := 0; j + i - 1 < n; j++ {
			k := j + i - 1
			if i == 2 {
				f[j][k] = s[j] == s[k]
			} else {
				f[j][k] = s[j] == s[k] && f[j + 1][k - 1]
			}
		}
	}

    
	dp := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := i; j >= 1; j-- {
			if f[j - 1][i - 1] {
				dp[i] = min(dp[i], dp[j - 1] + 1)
			}
		}
	}

	return dp[n] - 1
}

func min(a, b int) int {
	if a <b {
		return a
	}

	return b
}