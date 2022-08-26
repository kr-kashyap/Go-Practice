/*
    https://leetcode.com/problems/palindrome-linked-list/
    260822
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    ans := true
    temp := &ListNode{}
    temp = head
    var nums []int
    for ; temp != nil ; temp=temp.Next {
        nums = append(nums , temp.Val) 
    }
    for i :=0 ; ans != false && i < len(nums) / 2 ; i++ {
        if nums[i] != nums[len(nums)-i-1] {
            ans = false
        }
    }
    return ans
}