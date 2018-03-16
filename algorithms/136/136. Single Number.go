package algorithms

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		// n^n == 0
		// a^b^a^b^a == a
		res ^= n
	}
	return res
}

// timeout
//func singleNumber(nums []int) int {
//	set := make(map[int]struct{})
//	for _, v := range nums {
//		if _, has := set[v]; has {
//			delete(set, v)
//		} else {
//			set[v] = struct{}{}
//		}
//	}
//	for e := range set {
//		return e
//	}
//	return 0
//}
