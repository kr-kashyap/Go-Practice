/*
    https://leetcode.com/problems/fibonacci-number/
    280822
*/
func fib(n int) int {
    a := -1
    b := 1
    for i :=0 ; i < n+1 ; i++ {
        a , b = b , a + b
    }
    return b
}