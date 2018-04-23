package algorithms

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		bytes := []byte(v)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		s := string(bytes)
		m[s] = append(m[s], v)
	}
	var ss [][]string
	for e := range m {
		ss = append(ss, m[e])
	}
	return ss
}
