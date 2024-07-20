func searchInsert(nums []int, target int) int {
    regresar := 0
    for i, val := range nums{
        if target == val{
            return i
        } else {
            if  val < target{
                regresar = i + 1
            }
        }
    }
    return regresar
}