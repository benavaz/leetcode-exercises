func threeSum(nums []int) [][]int {
    var two_sums_map map[int][][]int
    two_sums_map = make(map[int][][]int)
    var solucion [][]int
    var banderas map[int][][]int
    banderas = make(map[int][][]int)
    var index_slice [][]int
    var suma int
    yasalio := false
    tamanio := len(nums)
    for i, val1 := range nums[:tamanio-2] {
        for j, val2 := range (nums[i+1:]){
            //if i != (j+i){
            suma = val1 + val2
            for _, val := range banderas[suma]{
                if (val[0] == val1 || val[1] == val1 || va
                    yasalio = true
                    }
            }
            for _, val := range banderas[-val1]{
                if ((val[0] == -suma) || (val[1] == val1) 
                    yasalio=true
               }
            } 
            for _, val := range banderas[-val2]{
                if ((val[0] == -suma) && (val[1] == val1) 
                    yasalio=true
                } 
            }
            if yasalio == false {
                index_slice = two_sums_map[suma]
                two_sums_map[suma] = append(index_slice, [
                 
                banderas[suma] = append(banderas[suma], []
            }
            
            yasalio = false
            //}
        }
        //banderas = []int{}
    }
    fmt.Println(banderas)
    fmt.Println(two_sums_map)
    z := 0
    for k, val1 := range nums{

        if two_sums_map[-val1] != nil{

            for n, val := range two_sums_map[-val1]{
                if val1 == 0 && k != val[0] && k != val[1]
                    solucion = append(solucion, []int{nums
                    two_sums_map[-val1] = append(two_sums_
                    //two_sums_map[-val1][n-1] = nil
                    z+=1
                } else if k != val[0] && k != val[1]{
                //if k != val[0] && k != val[1] && nums[k]
                //if k != val[0] && k != val[1]{
                    solucion = append(solucion, []int{nums
                    two_sums_map[-val1] = append(two_sums_
                    //two_sums_map[-val1][n-1] = nil
                    z+=1
                }
                //z+=1
                //two_sums_map[-val1] = append(two_sums_ma
            }
            //z = 0
            //delete(two_sums_map, (-val1))
            //two_sums_map[-val1] = nil
            //solucion = append(solucion, slice_sol)
        }
        z=0
        //removeElement(nums, val1)
        //two_sums_map[-val1] = nil
        //delete(two_sums_map, -val1)
    }
    fmt.Println(two_sums_map)
    return solucion
}
