package algorithms

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		} else {
			m[n] = struct{}{}
		}
	}
	return false
}
