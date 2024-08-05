func searchRange(nums []int, target int) []int {
    output := []int {-1, -1}
    for i, val := range nums{
        if val == target && output[0] == -1{
            output[0] = i
            output[1] = i
        } else if val == target {
            output[1] = i
        }
    }
    return output
}