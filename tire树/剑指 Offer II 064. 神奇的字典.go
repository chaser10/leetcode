package tire树

type tire map[rune]tire

type MagicDictionary struct {
	root tire
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{tire{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		cur := this.root
		for _, ch := range word {
			if cur[ch] == nil {
				cur[ch] = tire{}
			}

			cur = cur[ch]
		}

		cur['#'] = tire{}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return dfs(this.root, searchWord, false)
}

func dfs(root tire, word string, modified bool) bool {
	if len(word) == 0 {
		return modified && root['#'] != nil
	}

	ch := rune(word[0])
	if root[ch] != nil && dfs(root[ch], word[1:], modified) {
		return true
	} else if !modified {
		for i := 0; i < 26; i++ {
			if root[rune('a'+i)] != nil && rune('a'+i) != ch && dfs(root[rune('a'+i)], word[1:], true) {
				return true
			}
		}
	}

	return false
}
