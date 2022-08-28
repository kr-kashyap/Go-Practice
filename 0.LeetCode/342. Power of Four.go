/*
	https://leetcode.com/problems/power-of-four/
	260822
*/

import "math"

func isPowerOfFour(n int) bool {
    d := math.Log(float64(n)) / math.Log(4.0)
    ans := true
    i := int(d)
    if d - float64(i) != 0 {
        ans = false
    }
    return ans
}