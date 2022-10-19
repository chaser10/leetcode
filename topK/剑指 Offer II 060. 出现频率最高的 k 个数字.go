package topK

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	l := []int{}
	ret := []int{}
	for _, v := range m {
		l = append(l, v)
	}

	sort.Ints(l)
	frequent := l[len(l)-k]
	for k, v := range m {
		if v >= frequent {
			ret = append(ret, k)
		}
	}

	return ret
}
