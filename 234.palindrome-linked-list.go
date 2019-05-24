/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow := head
	fast := head
	var temp, newHead *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		temp = slow.Next
		slow.Next = newHead
		newHead = slow
		slow = temp
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != newHead.Val {
			return false
		}
		slow = slow.Next
		newHead = newHead.Next
	}
	return true
}