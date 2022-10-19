package 栈

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		num, err := strconv.ParseInt(token, 10, 32)
		if err == nil {
			stack = append(stack, int(num))
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if token == "+" {
				stack = append(stack, a+b)
			} else if token == "-" {
				stack = append(stack, a-b)
			} else if token == "*" {
				stack = append(stack, a*b)
			} else {
				stack = append(stack, a/b)
			}
		}
	}

	return stack[0]
}
