func containsDuplicate(nums []int) bool {
    var check_list []int
    for _, v := range (nums){
        for _, cv := range(check_list){
            if cv == v {
                return true
            }
        }
        check_list = append(check_list, v)
    }
    return false
}
