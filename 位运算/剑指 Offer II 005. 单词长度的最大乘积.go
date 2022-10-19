package 位运算

func maxProduct(words []string) int {
	mask := make([]int, len(words))
	countMask := func(word string) int {
		ret := 0
		for _, ch := range word {
			ret |= 1 << int(ch-'0')
		}

		return ret
	}

	for i, word := range words {
		mask[i] = countMask(word)
	}

	ans := 0
	for i := 0; i < len(mask); i++ {
		for j := i + 1; j < len(mask); j++ {
			if mask[i]&mask[j] == 0 {
				if len(words[i])*len(words[j]) > ans {
					ans = len(words[i]) * len(words[j])
				}
			}
		}
	}

	return ans
}
