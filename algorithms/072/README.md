# Edit Distance

## 描述

```txt

Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)



You have the following 3 operations permitted on a word:



a) Insert a character
b) Delete a character
c) Replace a character

```

## 题解

```go
package algorithms

//Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)
//
//You have the following 3 operations permitted on a word:
//
//a) Insert a character
//b) Delete a character
//c) Replace a character

//     a b c e
//   0 1 2 3 4
// c 1 1 2 2 3
// d 2 2 2 3 4
// e 3 3 3 3 ?
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i, v := range word1 {
		for j, vv := range word2 {
			switch vv {
			case v:
				dp[i+1][j+1] = dp[i][j]
			default:
				dp[i+1][j+1] = minOfThree(dp[i+1][j]+1, dp[i][j+1]+1, dp[i][j]+1)
			}
		}
	}
	return dp[m][n]
}

func minOfThree(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

```