func generateParenthesis(n int) []string {
	// solutions will be added to this slice of strings
	solution := []string{}

	var backtrack func(current_string string, open_count int, close_count int)
	backtrack = func(current_string string, open_count int, close_count int) {
		// if the current string is len 2*n, then it is a valid solution and it
		// is added to the result
		if len(current_string) == 2*n {
			solution = append(solution, current_string)
			return
		}

		// if we can add an open parenthese, we do so
		if open_count < n {
			backtrack(current_string+"(", open_count+1, close_count)
		}

		// if we can add a closing parentheses, we do so
		if close_count < open_count {
			backtrack(current_string+")", open_count, close_count+1)
		}
	}

	backtrack("", 0, 0)

	return solution
}