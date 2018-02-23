package algorithms

//You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
//For example, given:
//s: "barfoothefoobarman"
//words: ["foo", "bar"]
//
//You should return the indices: [0,9].
//(order does not matter).

// 找到只包含words的连续字字符串
// 解题思路是计数比较
func findSubstring(s string, words []string) []int {
	wordsMap := make(map[string]int)
	// 遍历words统计每个word的个数
	for _, v := range words {
		wordsMap[v] = wordsMap[v] + 1
	}
	wordLen := len(words[0])
	l := wordLen * len(words)
	e := len(s) - l + 1
	var result []int
	// 每次截取长度为word*word总数长度的字符串
	for i := 0; i < e; i++ {
		tempMap := make(map[string]int)
		// 截取字符串中的word计数
		var j int
		for ; j < len(words); j++ {
			index := i + j*wordLen
			word := s[index : index+wordLen]
			if _, ok := wordsMap[word]; ok {
				tempMap[word] = tempMap[word] + 1
				// 计数发现word长度，当出现个数超过，说明不符合条件，j<len(words)
				if tempMap[word] > wordsMap[word] {
					break
				}
			} else {
				break
			}
		}
		if j == len(words) {
			result = append(result, i)
		}
	}
	return result
}
