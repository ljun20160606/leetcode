package algorithms

func subsets(nums []int) [][]int {
	var result [][]int
	helper(&result, []int{}, nums, 0)
	return result
}

func helper(result *[][]int, temp []int, nums []int, index int) {
	t := make([]int, len(temp))
	copy(t, temp)
	*result = append(*result, t)
	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		helper(result, temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}
