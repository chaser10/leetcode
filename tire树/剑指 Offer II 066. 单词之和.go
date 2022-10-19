package tire树

type MapSum struct {
	node  map[rune]*MapSum
	count int
	m     map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{map[rune]*MapSum{}, 0, map[string]int{}}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this
	for _, ch := range key {
		if cur.node[ch] == nil {
			cur.node[ch] = &MapSum{map[rune]*MapSum{}, 0, map[string]int{}}
		}

		cur = cur.node[ch]
		cur.count += val - this.m[key]
	}

	this.m[key] = val

	cur.node['#'] = &MapSum{map[rune]*MapSum{}, 0, map[string]int{}}
}

func (this *MapSum) Sum(prefix string) int {
	cur := this
	for _, ch := range prefix {
		if cur.node[ch] != nil {
			cur = cur.node[ch]
		} else {
			return 0
		}
	}

	return cur.count
}
