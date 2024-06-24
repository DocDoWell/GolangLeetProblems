package spiralMatrix

func spiralOrder(matrix [][]int) []int {
	output := []int{}

	if len(matrix) == 0 {
		return output
	}

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			output = append(output, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			output = append(output, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				output = append(output, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				output = append(output, matrix[i][left])
			}
			left++
		}
	}

	return output
}