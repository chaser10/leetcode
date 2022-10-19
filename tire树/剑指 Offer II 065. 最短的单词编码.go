package tire树

type tire map[rune]tire

func minimumLengthEncoding(words []string) int {
	root := tire{}
	for _, word := range words {
		cur := root
		for i := len(word) - 1; i >= 0; i-- {
			ch := rune(word[i])
			if cur[ch] == nil {
				cur[ch] = tire{}
			}

			cur = cur[ch]
		}

		cur['#'] = tire{}
	}

	ret := 0
	var dfs func(root tire, cnt int)
	dfs = func(root tire, cnt int) {
		flag := false
		for i := 0; i < 26; i++ {
			if root[rune('a'+i)] != nil {
				flag = true
				dfs(root[rune('a'+i)], cnt+1)
			}
		}

		if !flag && root['#'] != nil {
			ret += cnt + 1
			return
		}
	}

	dfs(root, 0)
	return ret
}
