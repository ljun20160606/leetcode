package algorithms

//Write a function to find the longest common prefix string amongst an array of strings.

func longestCommonPrefix(strs []string) string {
	lenStrs := len(strs)
	if lenStrs < 1 {
		return ""
	}
	d := strs[0]
	lenMin := len(d)
	for i := 1; i < lenStrs; i++ {
		if lenMin > len(strs[i]) {
			lenMin = len(strs[i])
		}
	}
	var l int
LOOP:
	for i := 0; i < lenMin; i++ {
		b := d[i]
		for ii := 1; ii < lenStrs; ii++ {
			if strs[ii][i] != b {
				break LOOP
			}
		}
		l++
	}
	return string(d[:l])
}
