func findMin(nums []int) int {
    max := nums[0]
    for _, val := range nums{
        if val < max{
            max = val
        }
    }
    return max
}
