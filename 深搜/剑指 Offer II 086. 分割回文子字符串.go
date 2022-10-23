func partition(s string) [][]string {
	ret := [][]string{}
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

	var dfs(int i, splits []stirng) {
		if i == n {
			ret = append(ret, append([]int{}, splits...))
		}

		for j := 0; j < n; j++ {
			if f[i][j] {
				splits= append(splits, s[i : j + 1])
				dfs(j + 1, splits)
				splits = splits(:len(splits) - 1)
			}
		}
	}

	dfs(0, []string{})
	return ret
}