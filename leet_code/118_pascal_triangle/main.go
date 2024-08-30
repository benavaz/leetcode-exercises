package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	rows := make([][]int, 0)
	rows = append(rows, []int{1})
	// fmt.Println(rows)
	for i := 2; i <= numRows; i++ {
		var row []int
		upperRow := rows[i-2]
		// fmt.Println(upper)
		for j := 1; j <= i; j++ {
			if j == 1 {
				row = append(row, 1)
			} else if j == i {
				row = append(row, 1)
			} else {
				row = append(row, upperRow[j-2]+upperRow[j-1])
			}
		}
		rows = append(rows, row)
	}
	// fmt.Println(rows)
	return rows
}
