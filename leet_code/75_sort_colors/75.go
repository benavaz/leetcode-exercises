// sortColors solves the problem in O(n)
func sortColors(nums []int) []int{
	l := 0
	m := 0
	h := len(nums) - 1

	for m <= h{
		if nums[m] == 0{
			nums[l], nums[m] = nums[m], nums[l]
			l += 1
			m += 1
		} else if nums[m] == 2{
			nums[h], nums[m] = nums[m], nums[h]
			h -= 1
		}else if nums[m] == 1{
			m += 1
		}
	}

}

// sortColors2 solves the problem i O(2n)
func sortColors(nums []int) {
    var(
        white []int
        red []int
        blue []int
    )

    for _, num := range(nums){
        if num == 0{
            white = append(white, num)
        } else if num == 1{
            red = append(red, num)
        } else{
            blue = append(blue, num)
        }
    }
    //fmt.Println(white)
    //fmt.Println(red)
    //fmt.Println(blue)
    white = append(white, red...)
    white = append(white, blue...)
    //fmt.Println(white)

    for i, num := range(white) {
        nums[i] = num
    }
    //fmt.Println(nums)
}

