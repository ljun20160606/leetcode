package algorithms

//A message containing letters from A-Z is being encoded to numbers using the following mapping:
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//Given an encoded message containing digits, determine the total number of ways to decode it.
//
//For example,
//Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).
//
//The number of ways decoding "12" is 2.

// 12212
// 1 | 1
// 12 | 2
// 123 | 3
// 1221 | 5
// 12212 | 8
//
// 1345
// 1 | 1
// 13 | 2
// 134 | 2

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] > '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
