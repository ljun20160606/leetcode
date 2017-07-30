package origin

func KSum(nums []int, target, K int, result []int, results *[][]int) {
	if len(nums) < K || K < 2 {
		return
	}

	if K == 2 {
		// 2-sum
		left, right := 0, len(nums)-1
		for left < right {
			if temp := nums[left] + nums[right]; temp == target {
				*results = append(*results, append(result, nums[left], nums[right]))
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			} else if temp < target {
				left++
			} else {
				right--
			}
		}
	} else {
		length := len(nums)-K
		for i := 0; i <= length; i++ {
			if target < nums[i]*K || target > nums[len(nums)-1]*K {
				break
			}
			if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
				KSum(nums[i+1:], target-nums[i], K-1, append(result, nums[i]), results)
			}
		}
	}
}
