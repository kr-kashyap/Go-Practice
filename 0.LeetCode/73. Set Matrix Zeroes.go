/*
    https://leetcode.com/problems/set-matrix-zeroes/
    280822
*/
func setZeroes(matrix [][]int)  {
    rows := make([]int, 0)
    cols := make([]int, 0)
    for i := 0 ; i < len(matrix) ; i++ {
        for j := 0 ; j < len(matrix[0]) ; j++ {
            if matrix[i][j] == 0 {
                rows = append(rows , i)
                cols = append(cols , j)
            }    
        }
    }
    for _ , vr := range rows {
        //fmt.Printf("Row = %d\n",vr)
        for vc := 0 ; vc < len(matrix[0]) ; vc++ {
            matrix[vr][vc] = 0
        }
    }    
    for _ , vc := range cols {
        //fmt.Printf("Col = %d\n",vc) 
        for vr := 0 ; vr < len(matrix) ; vr++ {    
            matrix[vr][vc] = 0 
        }
    }
}