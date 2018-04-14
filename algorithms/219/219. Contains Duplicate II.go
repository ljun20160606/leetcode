package algorithms

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, has := m[v]; has {
			if i-vv <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}
