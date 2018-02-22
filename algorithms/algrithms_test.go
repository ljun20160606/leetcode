package algorithms

import (
	"reflect"
	"testing"
)

func TestArray2BinaryTree(t *testing.T) {
	tree := Array2BinaryTree([]interface{}{1, 2, 3})
	if tree.Val != 1 {
		t.Fail()
	}
	if tree.Left.Val != 2 {
		t.Fail()
	}
	if tree.Right.Val != 3 {
		t.Fail()
	}
}

func TestBreadthFirstBinaryTree(t *testing.T) {
	foo := []interface{}{1, 2, 2, 3, 4, 4, 3}
	root := Array2BinaryTree(foo)
	var bar []interface{}
	BreadthFirstBinaryTree(root, func(node *TreeNode) {
		bar = append(bar, node.Val)
	})
	if !reflect.DeepEqual(foo, bar) {
		t.Fail()
	}
}

func TestArray2ListNode(t *testing.T) {
	expected := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{Val: 5}}}}}
	result := Array2ListNode([]int{1, 2, 3, 4, 5})
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
}
