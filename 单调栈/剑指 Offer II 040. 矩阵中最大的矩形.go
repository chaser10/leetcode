package 单调栈

func maximalRectangle(matrix []string) int {
	levels := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		level := make([]int, len(matrix[0]))
		levels[i] = level
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 {
				if matrix[i][j] == '1' {
					levels[i][j] = 1
				}
			} else {
				if matrix[i][j] == '1' {
					levels[i][j] = levels[i-1][j] + 1
				} else {
					levels[i][j] = 0
				}
			}
		}
	}

	ret := 0
	for _, level := range levels {
		cur := largestRectangleArea(level)
		if cur > ret {
			ret = cur
		}
	}

	return ret
}
