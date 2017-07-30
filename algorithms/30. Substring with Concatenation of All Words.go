package algorithms

//You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
//For example, given:
//s: "barfoothefoobarman"
//words: ["foo", "bar"]
//
//You should return the indices: [0,9].
//(order does not matter).

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	wordsMap := make(map[string]int)
	for _, v := range words {
		wordsMap[v] = wordsMap[v] + 1
	}
	wordLen := len(words[0])
	l := wordLen * len(words)
	e := len(s) - l + 1
	result := []int{}
	for i := 0; i < e; i++ {
		tempMap := make(map[string]int)
		var j int
		for ; j < len(words); j++ {
			index := i + j * wordLen
			word := s[index:index+wordLen]
			if _, ok := wordsMap[word]; ok {
				tempMap[word] = tempMap[word] + 1
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
