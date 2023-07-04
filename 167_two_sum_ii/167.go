func twoSum(nums []int, target int) []int {
    l := 0
    r := len(nums) - 1
    sum := target + 1
    for sum != target{
        sum = nums[l] + nums[r]
        if sum < target {
            l += 1
        } else if sum > target{
            r -= 1
        }
    }
    return []int{l+1, r+1}
}
