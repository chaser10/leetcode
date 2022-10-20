func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i,j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := [][]int{}
	for _,interval := range intervals {
		if len(ret) == 0 {
			ret = append(ret, interval)
		} else {
			pre := ret[len(ret) - 1]
			if pre[1] < interval[0] {
				ret = append(ret, interval)
			} else {
				if interval[1] >= pre[1] {
					pre[1] = interval[1]
				}
			}
		}
	}

	return ret
}