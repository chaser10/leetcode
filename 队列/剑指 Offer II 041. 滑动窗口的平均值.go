package 队列

type MovingAverage struct {
	window []int
	sum    float64
	size   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		sum:    0.0,
		window: []int{},
		size:   size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.window) < this.size {
		this.sum += float64(val)
		this.size += 1
	} else {
		this.sum -= float64(this.window[0])
		this.window = this.window[1:]
		this.sum += float64(val)
	}

	this.window = append(this.window, val)
	return this.sum / float64(len(this.window))
}
