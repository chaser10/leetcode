func lenLongestFibSubseq(arr []int) int {
	ret := 0
	n := len(arr)
	m := map[int]int{}
    f := make([][]int, n)
	for i,num := range arr {
		m[num] = i
        f[i] = make([]int, n)
	}

	
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			pre := arr[j] - arr[i]
			if idx,ok := m[pre]; ok && idx < i{
				f[i][j] = max(f[idx][i] + 1, 3)
				ret = max(ret, f[i][j])
			}
		}

	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}