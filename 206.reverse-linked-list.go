/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// iter
func reverseList(head *ListNode) *ListNode {
	var newHead, temp *ListNode
	for head != nil {
		temp = head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}

//recur
// func reverseList(head *ListNode) *ListNode {
// 	return recur(head, nil)
// }

// func recur(head, prev *ListNode) *ListNode {
// 	if head == nil {
// 		return prev
// 	}
// 	temp := head.Next
// 	head.Next = prev
// 	return recur(temp, head)
// }

