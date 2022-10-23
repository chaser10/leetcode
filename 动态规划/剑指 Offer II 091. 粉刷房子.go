func minCost(costs [][]int) int {
	n = len(costs)
	f := make([][]int, n + 1)
	dp := make([]int, n + 1)
	for i := range f {
		f[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		f[1][i] = costs[0][i]
	}

	dp[1] = min(costs[0][0], min(costs[0][1], costs[0][2]))
	for i := 2; i < n; i++ {
		for j := 0; j < 3; j++ {
			if j == 0 {
				f[i][j] = dp[i - 2] + min(f[i - 1][1], f[i - 1][2])
			} else if j == 1 {
				f[i][j] = dp[i - 2] + min(f[i - 1][0], f[i - 1][2])
			} else {
				f[i][j] = dp[i - 2] + min(f[i - 1][1], f[i - 1][0])
			}
		}

		dp[i] = min(f[i][0], min(f[i][1], f[i][2]))
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}