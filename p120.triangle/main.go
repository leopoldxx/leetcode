package main
func min(a, b int) int{
    if a < b {
        return a
    }
    return b
}


func minimumTotal(triangle [][]int) int {
    if len(triangle) == 1 {
        return triangle[0][0]
    }
    for i := 1; i < len(triangle); i++ {
        for j:= 0; j <= i; j++ {
            if j == 0 {
                triangle[i][j] = triangle[i][j] +  triangle[i-1][j]
            } else if j == i {
                triangle[i][j] = triangle[i][j] +  triangle[i-1][j-1]
            } else {
                triangle[i][j] = triangle[i][j] + min(triangle[i-1][j-1], triangle[i-1][j])
            }
        }
    }
    last := triangle[len(triangle)-1]
    minSum := last[0]
    for i := 1; i < len(last); i++ {
        if minSum > last[i] {
            minSum = last[i]
        }
    }
    return minSum
}