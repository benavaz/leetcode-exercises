func trap(nums []int) int {
	var water int
	var tmp int
	indexes := make(map[int][]int)
	for i, v := range nums {
	  for j := v; j >= 0; j-- {
		  if len(indexes[j]) > 0{ 
			tmp += i - indexes[j][0] -1
		  //  indexes[j][0] = i
		  } 
		  indexes[j] = append(indexes[j], i)
		
	  }
	  //fmt.Println(tmp)
	  water += countWater(indexes)
	  //fmt.Println(indexes)
	}
	return water
  }
func countWater(walls map[int][]int) int {
	var water, tmp int
	//fmt.Println(walls)
	for i, index := range walls {
	  tmp = 0
	  for i := 0; i < len(index)-1; i++ {
		tmp = index[i+1] - index[i] - 1
		water += tmp
		index =index[1:]
	}
	walls[i] = index
	//water += tmp
  }
  //fmt.Println(walls)
  return water
}