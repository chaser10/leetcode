package 队列

type RecentCounter struct {
	stack []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.stack) > 0 && t-this.stack[0] > 3000 {
		this.stack = this.stack[1:]
	}

	this.stack = append(this.stack, t)
	return len(this.stack)
}
