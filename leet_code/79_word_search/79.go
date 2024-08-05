

func exist(board [][]byte, word string) bool {
	y := len(board)
	x := len(board[0])

	var count int
	//tmp_word := ""

	var node [][]int

	for j := 0; j < y ; j++ {
		for i := 0; i < x ; i++ {
			if board[j][i] == word[0]{
				node = append(node, []int{j,i})
				count count += 1

			}
		}
	}

}

func findSiblings (x, y, i, j int) ([][]int){
	if i > 0 && j >0{
		return []int{}
	}
}

func traverseGraph(board [][]byte){
	y := len(board) - 1
	x := len(board[0]) - 1
	var j, i int
	var letter string

	var node [][]int

	node = append(node, []int{j, i})
	for len(node) > 0 {
		i = node[0][1]
		j = node[0][0]
		if i < x && j < y{
			node = append(node, []int{j, i+1})
			node = append(node, []int{j+1, i})
		} else if i < x {
			node = append(node, []int{j, i+1})
		} else if j < y {
			node = append(node, []int{j+1, i})
		}
		node = node[0:]
		fmt.Println(board[i][j])
	}

	fmt.Println(letter)
}