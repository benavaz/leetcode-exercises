func numSquares(n int) int {
	// Use a proper queue structure for BFS
	nodes := []node{}
	visited := make([]bool, n+1) // Fix the visited array size
	squares := generatePerfectSquares(n)

	// Start BFS
	nodes = append(nodes, node{n, 0})
	visited[n] = true

	for len(nodes) > 0 {
		// Get the first node (queue-like behavior)
		current_node := nodes[0]
		nodes = nodes[1:] // Dequeue
		steps := current_node.steps

		// Explore all possible next nodes by subtracting perfect squares
		for _, square := range squares {
			next_val := current_node.n - square
			if next_val == 0 {
				return steps + 1
			}
			if next_val > 0 && !visited[next_val] {
				visited[next_val] = true
				nodes = append(nodes, node{next_val, steps + 1}) // Proper BFS enqueue
			}
		}
	}
	return -1
}

// node represents a node in the BFS
type node struct {
	n     int // current node number
	steps int // number of steps to that number
}

// function to generate the squares < n
func generatePerfectSquares(n int) []int {
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}
	return squares
}