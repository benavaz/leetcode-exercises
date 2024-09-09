func permute(nums []int) [][]int {
	var output [][]int
	current_permutation := []int{}
	valid_size := len(nums)

	var backtrack func(nums []int, current_permutation []int)

	backtrack = func(nums []int, current_permutation []int) {
		if len(current_permutation) == valid_size {
			output = append(output, current_permutation)
			return
		}

		current_permutation_copy := make([]int, len(current_permutation))
		copy(current_permutation_copy, current_permutation)
		for i := 0; i < len(nums); i++ {
			current_permutation = append(current_permutation_copy, nums[i])
			nums_copy := make([]int, len(nums))
			copy(nums_copy, nums)
			nums_copy = append(nums_copy[:i], nums_copy[i+1:]...)

			backtrack(nums_copy, current_permutation)
			current_permutation = []int{}
		}

	}

	backtrack(nums, current_permutation)

	return output
}