func asteroidCollision(asteroids []int) []int {
    output := []int{}
    var cuenta int
    var vo int

    for _, v := range(asteroids){
        if len(output) == 0{
            output = append(output, v)
        } else if output[len(output) - 1] < 0 && v < 0{
            output = append(output, v)
        }else if v > 0 {
            output = append(output, v)
        }else if output[len(output) - 1] > 0 && v < 0{
            cuenta = 0
            for i := len(output)-1; i >= 0; i-- {
                vo = output[i]
                if vo > 0 && vo > v*-1{   
                    output = output[:len(output) - cuenta]
                    break
                } else if vo == v*-1{
                    cuenta += 1
                    output = output[:len(output) - cuenta]
                    break 
                } else if vo > 0 && vo < v*-1{
    
                    cuenta += 1
                }

                if i == 0 && vo < v*-1{
                    output = output[:len(output) - cuenta]
                    output = append(output, v)
                    break
                }
            }
        }
    }
    return output 
}
