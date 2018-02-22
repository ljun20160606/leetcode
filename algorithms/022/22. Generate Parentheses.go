package algorithms

//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//For example, given n = 3, a algorithms set is:
//
//[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]

// 卡特兰数 https://zh.wikipedia.org/wiki/%E5%8D%A1%E5%A1%94%E5%85%B0%E6%95%B0
func generateParenthesis(n int) []string {
	var result []string
	chars := make([]byte, 0, n*2)
	gp(&result, chars, n, n)
	return result
}

func gp(result *[]string, chars []byte, left, right int) {
	if left|right == 0 {
		*result = append(*result, string(chars))
		return
	}
	if left > 0 {
		gp(result, append(chars, '('), left-1, right)
	}
	if left < right {
		gp(result, append(chars, ')'), left, right-1)
	}
}
