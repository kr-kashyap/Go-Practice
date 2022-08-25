/*
    https://leetcode.com/problems/ransom-note/
    250822
*/

//import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
    i:=0
    ans:=false
    magmap := make(map[string]int)
    for i=0; i<len(magazine); i++ {
        if _, ok := magmap[string(magazine[i])]; ok == true {
            magmap[string(magazine[i])] += 1 
        } else {
            magmap[string(magazine[i])] = 1
        }
    }
    
    rnmap := make(map[string]int)
    for i=0; i<len(ransomNote); i++ {
        if _, ok := rnmap[string(ransomNote[i])]; ok == true {
            rnmap[string(ransomNote[i])] += 1 
        } else {
            rnmap[string(ransomNote[i])] = 1
        }
    }
    
    for key, element := range rnmap {
        //fmt.Printf("%s %d %d\n", key, magmap[key], element)
        if magmap[key] < element {
            //fmt.Println("FALSE\n")
            ans = false
            return ans
        } else {
            ans = true
        }
    } 
    //fmt.Println()
    return ans
}