/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	solution := [][]int{}
	queue := []*TreeNode{}
	tmp := []int{}
	tmp_target := 1
	var element *TreeNode

	if root == nil {
		return solution
	}

	// first enqueue first element outside loop
	queue = append(queue, root)
	// while queue is not empty
	for len(queue) != 0 {

		element = queue[0]
		tmp = append(tmp, element.Val)
		queue = queue[1:]

		// enqueue left and right nodes for each node
		if element.Left != nil {
			queue = append(queue, element.Left)
		}
		if element.Right != nil {
			queue = append(queue, element.Right)
		}

		tmp_target--

		if tmp_target == 0 {
			solution = append(solution, tmp)
			tmp = nil
			tmp_target = len(queue)
		}

	}

	return solution
}