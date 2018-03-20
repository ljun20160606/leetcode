package algorithms

func subsets(nums []int) [][]int {
	var result [][]int
	helper(nums, &result, []int{}, 0)
	return result
}

func helper(nums []int, result *[][]int, temp []int, index int) {
	t := make([]int, len(temp))
	copy(t, temp)
	*result = append(*result, t)
	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		helper(nums, result, temp, i+1)
		temp = temp[:len(temp)-1]
	}
}
