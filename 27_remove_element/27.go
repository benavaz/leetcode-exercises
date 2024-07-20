func removeElement(nums []int, val int) int {
    k := 0
    diferentes := []int{}
    for _, num := range nums {
        if num == val{
        } else {
            k++
            diferentes = append(diferentes, num)
        }
    }
    for i, difs := range diferentes{
        nums[i] = difs
    }
    return k
}