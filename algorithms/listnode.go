package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

func Array2ListNode(array []int) *ListNode {
	head := new(*ListNode)
	current := head
	for e := range array {
		if *current == nil {
			*current = new(ListNode)
		}
		(**current).Val = array[e]
		current = &((*current).Next)
	}
	return *head
}
