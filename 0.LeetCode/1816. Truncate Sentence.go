/*
    https://leetcode.com/problems/truncate-sentence/
    260822
*/
func truncateSentence(s string, k int) string {
    i := 0
    for ; k>0 && i < len(s) ; i++ {
        if s[i] == ' ' {
            k--
        }
    }
    s1 := s[ : 0]
    if i < len(s) {
        s1 = s[ : i - 1]    
    } else {
        s1 = s[ : i]    
    }
    return s1
}