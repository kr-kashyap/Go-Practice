/*
    https://leetcode.com/problems/middle-of-the-linked-list/
    250822
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
    temp := &ListNode{}
    temp = head
    len:=0
    for ; temp != nil ; temp=temp.Next {
        len++
    }
    temp = head
    for i:=0 ; i < len/2 ; i++   {
        temp=temp.Next
    }
    return temp
}