package tire树

type Trie struct {
	child [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{child: [26]*Trie{}, isEnd: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this
	for i := 0; i < len(word); i++ {
		if p.child[word[i]-'a'] == nil {
			p.child[word[i]-'a'] = &Trie{child: [26]*Trie{}, isEnd: false}
		}

		p = p.child[word[i]-'a']
	}

	p.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this
	for i := 0; i < len(word); i++ {
		u := word[i] - 'a'
		if p.child[u] == nil {
			return false
		}

		p = p.child[u]
	}

	return p.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for _, ch := range prefix {
		if p.child[ch-'a'] == nil {
			return false
		}

		p = p.child[ch-'a']
	}

	return true
}
