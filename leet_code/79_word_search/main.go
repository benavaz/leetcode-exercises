func exist(board [][]byte, word string) bool {
	var backtrack func(i int, j int, word string)
	// variable used to append the letters that already appeared
	//var seen [][]int
	exists := false
	y := len(board)
	x := len(board[0])

	//fmt.Println(y)
	//fmt.Println(x)
	//fmt.Println(seen)

	//seen = append(seen, []int{45,67})
	//fmt.Println(seen)
	//fmt.Println(contains(seen, []int{1,1}))
	//fmt.Println(contains(seen, []int{45,67}))
	backtrack = func(j int, i int, word string) {
		// if letter [i][j] != word[i]
		//fmt.Println(word)
		//letter := board[j][i]
		//fmt.Printf("%d, %d ",i,j)

		// if the letter does not match or was already chosen
		//if letter != word[0] {
		//fmt.Println("letter does not match")
		//    return
		//}

		if len(word) == 0 {
			exists = true
			return
		}

		//fmt.Println("match")

		if j < 0 || j >= y || i < 0 || i >= x || board[j][i] != word[0] {
			return
		}

		// Temporarily mark the current cell as visited
		temp := board[j][i]
		board[j][i] = '#'

		// Recursively explore the four directions (down, up, right, left)
		backtrack(j+1, i, word[1:])
		backtrack(j-1, i, word[1:])
		backtrack(j, i+1, word[1:])
		backtrack(j, i-1, word[1:])

		board[j][i] = temp

	}

	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			//fmt.Println("---------")
			//fmt.Println(j)
			//fmt.Println(i)
			//fmt.Println("---------")
			//backtrack(j, i, [][]int{}, word)
			//if exists {
			//    break
			//}
			if board[j][i] == word[0] { // Start search only from matching first letter
				backtrack(j, i, word)
				if exists {
					return true
				}
			}
		}
	}
	// backtrack(0, 0, seen, word)
	//return exists
	//backtrack(0,0,[][]int{},word)
	return exists
}

func contains(slice [][]int, element []int) bool {
	for _, subSlice := range slice {
		if reflect.DeepEqual(subSlice, element) {
			return true
		}
	}
	return false
}