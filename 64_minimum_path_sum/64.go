func minPathSum(grid [][]int) int {
    y := len(grid)  
    x := len(grid[0])  
    col_vals := make([]int, x)
    row_vals := make([]int, y)
    var val int
 
    if x == 1 && y == 1{
        return grid[0][0]
    }
    for j :=0; j < y ; j++{
        for i := 0; i < x; i++{
            val = grid[j][i]
            if i > 0 && j > 0{
                col_vals[i] = min(row_vals[j], col_vals[i]) + val
                row_vals[j] = col_vals[i]
            } else if j > 0 {
                row_vals[j] = val + col_vals[i]
                col_vals[i] = row_vals[j]
            } else if i > 0 {
                col_vals[i] = val + row_vals[j]
                row_vals[j] = col_vals[i]
            } else if i == 0 && j==0{
                row_vals[j] = val
                col_vals[i] = val
            } 
        }
    }
    return row_vals[y-1]
}

func min(a ,b int) int{
    if a < b {
        return a
    }
    return b
}
