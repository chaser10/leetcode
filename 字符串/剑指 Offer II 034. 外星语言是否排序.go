package 字符串

func isAlienSorted(words []string, order string) bool {
	m := map[int]int{}
	for i, ch := range order {
		m[int(ch)] = i
	}

	for i := 0; i < len(words)-1; i++ {
		j := 0
		prefixEqual := true
		for j < len(words[i]) && j < len(words[i+1]) {
			if words[i][j] != words[i+1][j] {
				if m[int(words[i][j])] > m[int(words[i+1][j])] {
					return false
				} else if m[int(words[i][j])] < m[int(words[i+1][j])] {
					prefixEqual = false
					break
				}
			}

			j++
		}

		if prefixEqual {
			if len(words[i]) > len(words[i+1]) {
				return false
			}
		}
	}

	return true
}
