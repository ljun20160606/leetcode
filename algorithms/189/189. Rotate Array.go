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
