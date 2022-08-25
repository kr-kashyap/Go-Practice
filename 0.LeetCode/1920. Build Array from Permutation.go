/*
	https://leetcode.com/problems/build-array-from-permutation/submissions/
	250822
*/

import "fmt"

func buildArray(nums []int) []int {
    var num = make([]int, len(nums))
    for i:=0;i<len(nums);i++ {  
        //fmt.Printf("%d %d\n",i,nums[nums[i]])
        num[i] = nums[nums[i]]    
    }
    return num
}