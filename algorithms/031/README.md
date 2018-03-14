# Next Permutation

## 描述

```txt

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.


If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).


The replacement must be in-place, do not allocate extra memory.


Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

```

## 题解

```go
package algorithms

//Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
//
//If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).
//
//The replacement must be in-place, do not allocate extra memory.
//
//Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1

// 找到当前数字的下一行
// 1, 2, 3, 4
// 1, 2, 4, 3
// 1, 3, 2, 4
// 下一行比上一行大，同时是最小可能
func nextPermutation(nums []int) {
	// 边界条件
	if len(nums) < 2 {
		return
	}
	i := len(nums) - 1
	// 向前遍历
	// 1, 2, 4, 3, 2
	for ; i > 0; i-- {
		// 2 < 4
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i != 0 {
		// 4, 3, 2 从右向左遍历找到第一个比2大的数字, 4左边的2和3交换位置
		for t := len(nums) - 1; t >= i; t-- {
			if nums[t] > nums[i-1] {
				nums[t], nums[i-1] = nums[i-1], nums[t]
				// 1, 3, 4, 2, 2
				// 此时4, 2, 2一定是降序，反转得到 1, 3, 2, 2, 4
				break
			}
		}
	}
	reverseSlice(nums[i:])
}

// 时间复杂度O(n)
func reverseSlice(nums []int) {
	if len(nums) < 2 {
		return
	}
	end := len(nums) - 1
	nums[0], nums[end] = nums[end], nums[0]
	reverseSlice(nums[1:end])
}

func nextPermutationBest(nums []int) {
	n := len(nums)
	i := n - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	if i != 0 {
		j := i
		for ; j < n-1; j++ {
			if nums[j+1] <= nums[i-1] {
				break
			}
		}
		nums[i-1], nums[j] = nums[j], nums[i-1]
	}
	for j, k := i, n-1; j < k; j++ {
		nums[j], nums[k] = nums[k], nums[j]
		k--
	}
}

```