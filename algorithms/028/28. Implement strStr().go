package algorithms

//Implement strStr().
//
//Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

func strStr(haystack, needle string) int {
	ls, lss := len(haystack), len(needle)
	if ls|lss == 0 {
		return 0
	}
	end := ls - lss + 1
	for i := 0; i < end; i++ {
		var t int
		for ii := range needle {
			if haystack[i+ii] != needle[ii] {
				break
			}
			t++
		}
		if t == lss {
			return i
		}
	}
	return -1
}
