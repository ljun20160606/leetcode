package algorithms

import (
	"fmt"
	"github.com/LFZJun/solution-lib/solution/origin"
	"github.com/davecgh/solution-spew/spew"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(twoSum([]int{1, 3, 2}, 3))
}

func Test1V1(t *testing.T) {
	fmt.Println(twoSumV0([]int{1, 2, 3}, 3))
}

func Test2(t *testing.T) {
	result := addTwoNumbers(
		&ListNode{5, nil},
		&ListNode{5, nil},
	)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func Test3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func TestFindKth(t *testing.T) {
	fmt.Println(findKth([]int{1, 2, 3, 4}, []int{2, 3, 4, 7, 9}, 3))
}

func Test4(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{3}))
}

func TestSym(t *testing.T) {
	fmt.Println(sym("abbbc"))
}

func Test5(t *testing.T) {
	fmt.Println(longestPalindrome("aba"))
}

func Test6(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

func Test7(t *testing.T) {
	fmt.Println(reverse(2147483648))
}

func Test8(t *testing.T) {
	fmt.Println(myAtoi("-2147483648"))
}

func Test9(t *testing.T) {
	fmt.Println(isPalindrome(121))
}

func Test10(t *testing.T) {
	fmt.Println(isMatch("ab", ".*c"))
}

func Test11(t *testing.T) {
	fmt.Println(maxArea([]int{1, 1}))
}

func Test12(t *testing.T) {
	fmt.Println(intToRoman(112))
}

func Test13(t *testing.T) {
	fmt.Println(romanToInt("IV"))
}

func Test14(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"a", "a", "a"}))
}

func Test15(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func Test16(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func Test17(t *testing.T) {
	fmt.Println(letterCombinations("0234"))
}

func Test18(t *testing.T) {
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
}

func TestKsum(t *testing.T) {
	var result [][]int
	origin.KSum([]int{-3, -1, 0, 2, 4, 5}, 1, 2, []int{}, &result)
	fmt.Println(result)
}

func Test19(t *testing.T) {
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, nil}}, 2))
}

func Test20(t *testing.T) {
	fmt.Println(isValid("(())"))
}

func Test21(t *testing.T) {
	spew.Println(mergeTwoLists(&ListNode{2, nil}, &ListNode{1, nil}))
}

func Test23(t *testing.T) {
	spew.Println(mergeKLists([]*ListNode{{2, nil}, {1, nil}}))
}

func Test26(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func Test27(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func Test28(t *testing.T) {
	fmt.Println(strStr("a", "a"))
}

func Test72(t *testing.T) {
	fmt.Println(minDistance("abce", "cde"))
}
