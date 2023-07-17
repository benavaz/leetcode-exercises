package main

import "fmt"

func main(){
    input := []int{0,2,0,1,1,0,0,2,0,1}
	fmt.Println(input)
    fmt.Println(sortColors(input))
}

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

    return nums
}
