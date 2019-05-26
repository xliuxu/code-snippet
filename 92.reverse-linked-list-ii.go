/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	newHead := &ListNode{}
	newHead.Next = head
	pre := newHead
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < (n - m); i++ {
		temp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = temp
	}
	return newHead.Next
}

