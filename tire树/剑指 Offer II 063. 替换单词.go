package tire树

import "strings"

type tire struct {
	child [26]*tire
	isEnd bool
}

func (this *tire) insert(word string) {
	p := this
	for _, ch := range word {
		if p.child[ch-'a'] == nil {
			p.child[ch-'a'] = &tire{[26]*tire{}, false}
		}

		p = p.child[ch-'a']
	}

	p.isEnd = true
}

func (this *tire) searchPrefix(word string) string {
	p := this
	cur := strings.Builder{}
	for _, ch := range word {
		u := ch - 'a'
		if p.child[u] == nil {
			return ""
		}

		cur.WriteRune(ch)
		p = p.child[u]
		if p.isEnd == true {
			return cur.String()
		}
	}

	return ""
}

func replaceWords(dictionary []string, sentence string) string {
	tireTree := tire{isEnd: false, child: [26]*tire{}}
	for _, word := range dictionary {
		tireTree.insert(word)
	}

	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		prefix := tireTree.searchPrefix(words[i])
		if prefix != "" {
			words[i] = prefix
		}
	}

	return strings.Join(words, " ")
}
