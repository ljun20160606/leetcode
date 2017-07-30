package algorithms

//Reverse a singly linked list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var newHead, tempHead *ListNode
	for head != nil {
		tempHead = head.Next
		head.Next = newHead
		newHead = head
		head = tempHead
	}
	return newHead
}
