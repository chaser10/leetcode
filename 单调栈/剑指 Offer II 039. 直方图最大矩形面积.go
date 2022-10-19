package 单调栈

import "fmt"

func largestRectangleArea(heights []int) int {
	left, right := []int{}, []int{}
	stack := []int{}
	for i, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > height {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			left = append(left, stack[len(stack)-1])
		} else {
			left = append(left, -1)
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			right = append(right, stack[len(stack)-1])
		} else {
			right = append(right, len(heights))
		}
		stack = append(stack, i)
	}

	ret := 0
	fmt.Println(left)
	fmt.Println(right)
	for i := 0; i < len(heights); i++ {
		if ret < (right[i]-left[i]-1)*heights[i] {
			ret = (right[i] - left[i] - 1) * heights[i]

		}
	}

	return ret
}
