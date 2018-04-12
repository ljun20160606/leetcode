# Rotate Array

## 描述

```txt
Rotate an array of n elements to the right by k steps.

For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].

Note:
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.

[show hint]

Hint:
Could you do it in-place with O(1) extra space?

Related problem: Reverse Words in a String II

Credits:
Special thanks to @Freezen for adding this problem and creating all test cases.

```

## 题解

```go
package algorithms

func rotate(nums []int, k int) {
	if k == 0 || k == len(nums) {
		return
	}
	if k > len(nums) {
		k %= len(nums)
	}
	reverse(nums[0 : len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
}

func reverse(s []int) {
	if len(s) == 0 {
		return
	}
	end := len(s) - 1
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		s[i], s[end-i] = s[end-i], s[i]
	}
}

```