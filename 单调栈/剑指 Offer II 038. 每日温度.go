package 单调栈

func dailyTemperatures(temperatures []int) []int {
	reverseSlice := func(sl []int) {
		l, r := 0, len(sl)-1
		for l < r {
			t := sl[r]
			sl[r] = sl[l]
			sl[l] = t
			l++
			r--
		}
	}

	reverseSlice(temperatures)
	stack1 := []int{}
	stack2 := []int{}
	ret := []int{}
	for i := 0; i < len(temperatures); i++ {
		for len(stack1) > 0 && stack1[len(stack1)-1] < temperatures[i] {
			stack1 = stack1[:len(stack1)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		if len(stack1) == 0 {
			ret = append(ret, 0)
		} else {
			ret = append(ret, i-stack2[len(stack2)-1])
		}
		stack1 = append(stack1, temperatures[i])
		stack2 = append(stack2, i)

	}

	reverseSlice(ret)
	return ret
}
