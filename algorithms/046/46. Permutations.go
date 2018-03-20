package algorithms

func permute(nums []int) [][]int {
	var result [][]int
	helper(nums, &result, []int{})
	return result
}

func helper(nums []int, result *[][]int, temp []int) {
	if len(nums) == len(temp) {
		t := make([]int, len(temp))
		copy(t, temp)
		*result = append(*result, t)
	}
Loop:
	for i := range nums {
		for e := range temp {
			if temp[e] == nums[i] {
				continue Loop
			}
		}
		temp = append(temp, nums[i])
		helper(nums, result, temp)
		temp = temp[:len(temp)-1]
	}
}
