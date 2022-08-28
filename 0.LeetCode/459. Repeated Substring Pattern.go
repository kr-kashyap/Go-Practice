/*
    https://leetcode.com/problems/repeated-substring-pattern/
    280822
*/

/*
    Logic:
        -Find the factorials of len(s).
        -For each factorial, try concatenating the substring of len fact[i] and check whether concatenated string results into the original string.
*/
func repeatedSubstringPattern(s string) bool {
    n := len(s)
    if (n == 1) {
        return false
    }
    ans := false
    fact := make([]int , 0)
    for i := 2 ; i <= n / 2 ; i++ {
        if n % i == 0 {
            fact = append(fact , i)
        }
    }
    if len(fact) == 0 {
        temp := ""
        for i := n ; i>0 ; i-- {
            temp += s [ 0 : 1 ]
        } 
        //fmt.Println(temp)
        if (temp == s) {
            return true
        }
    }
    for _ , v := range fact {
        //fmt.Println(v)
        temp := ""
        for i := n / v ; i>0 ; i-- {
            temp += s [ 0 : v ]
        } 
        //fmt.Println(temp)
        if (temp == s) {
            return true
        }
    }
    return ans
}