func maxSubArray(nums []int) int {
    tamanio := len(nums)
    izq_tmp := 0
    der_tmp := 0
    derecha := make([]int, tamanio)
    izquierda := make([]int, tamanio)
    for i, _ := range nums {
        if izq_tmp + nums[i] <= 0{
            izq_tmp = 0
            izquierda[i] = 0  
        } else {
            izq_tmp += nums[i]
            izquierda[i] = izq_tmp
        }
        if der_tmp + nums[tamanio - i - 1] <= 0 {
            der_tmp = 0
        } else {
            der_tmp += nums[tamanio - i - 1]
            derecha[tamanio - i - 1] = der_tmp
        }
    }
    max := 0
    max_num := nums[0]
    //var index_max int
    var resultado int
    for i, _ := range izquierda {
        //if izquierda[i] == 0 && derecha[i] == 0{
        //    resultado = izquierda[i]
        //}else{
        resultado = izquierda[i] + derecha[i] 
        if resultado == 0 {
            resultado = izquierda[i]
        } else{
            resultado += (nums[i] * -1)
        }
        //}
        if resultado > max {
            max = resultado
            //index_max = i
        }
        if nums[i] > max_num{
            max_num = nums[i]
        }
    }
    fmt.Println(izquierda)
    fmt.Println(derecha)
    fmt.Println(max) 
    fmt.Println(max_num)

    if max == 0{
        return max_num
    } else {
    return max
    }
}

