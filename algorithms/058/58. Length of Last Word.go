package algorithms

func lengthOfLastWord(s string) int {
	var r int
	for e := range s {
		if s[e] != ' ' {
			if e > 0 && s[e-1] == ' ' {
				r = 0
			}
			r++
		}
	}
	return r
}
