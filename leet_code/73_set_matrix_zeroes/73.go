// ineficient solution

func setZeroes(matrix [][]int) {
	output := make([][]int, len(matrix))
	deepCopy(matrix, output)
	for i, row := range matrix {
		// fmt.Println(val)
		for j, val := range row {
			// fmt.Println(val)
			if val == 0 {
				// set row to zeroes
				for a, _ := range output[i] {
					output[i][a] = 0
				}

				// set column to zeroes
				for _, row := range output {
					for b, _ := range row {
						if b == j {
							row[b] = 0
						}
					}
				}
			}
		}
	}
	// fmt.Println(matrix)
	// fmt.Println(output)
	deepCopy(output, matrix)
}

// deepCopy creates a deepCopy (element by element)}
// of one matrix to dst matrix
func deepCopy(matrix [][]int, dst [][]int) {
	// output := make([][]int, len(matrix))
	for i := range matrix {
		inner := make([]int, len(matrix[i]))
		copy(inner, matrix[i])
		dst[i] = inner
	}

}