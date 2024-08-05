func maxSubArray(nums []int) int {
    sum := 0
    max_sum := 0
    for i, val := range nums {
        if sum < 0{
            sum = val
        } else{
            sum += val
        }
        if i == 0{
            max_sum = val
        } else if sum > max_sum {
            max_sum = sum
        }
    }
    return max_sum
}
