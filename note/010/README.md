# Regular Expression Matching 

## 描述

    Implement regular expression matching with support for '.' and '*'.

    '.' Matches any single character.
    '*' Matches zero or more of the preceding element.

    The matching should cover the entire input string (not partial).

    The function prototype should be:
    bool isMatchV0(const char *s, const char *p)

    Some examples:
    isMatchV0("aa","a") ? false
    isMatchV0("aa","aa") ? true
    isMatchV0("aaa","aa") ? false
    isMatchV0("aa", "a*") ? true
    isMatchV0("aa", ".*") ? true
    isMatchV0("ab", ".*") ? true
    isMatchV0("aab", "c*a*b") ? true

## 题解

### 1. 动态规划

#### 图

|     |       | a     | *     | c     |
| --- | ----- | ----- | ----- | ----- |
|     | true  | false | true  | false |
| a   | false | true  | true  | false |
| b   | false | false | false | false |
| c   | false | false | false | false |

```go
func isMatch(s string, p string) bool {
    lenY, lenX := len(s), len(p)
    dp := make([][]bool, lenY+1)
    for y := range dp {
        dp[y] = make([]bool, lenX+1)
    }
    dp[0][0] = true
    for x := 1; x <= lenX; x++ {
        if p[x-1] == '*' {
            dp[0][x] = dp[0][x-2]
        }
    }
    for y, mv := range s {
        for x, pv := range p {
            switch pv {
            case '*':
                if p[x-1] == s[y] || p[x-1] == '.' {
                    dp[y+1][x+1] = dp[y][x+1] || dp[y+1][x] || dp[y+1][x-1]
                } else {
                    dp[y+1][x+1] = dp[y+1][x-1]
                }
            case '.', mv:
                dp[y+1][x+1] = dp[y][x]
            }
        }
    }
    return dp[lenY][lenX]
}
```

### 2. 递归、回溯

```go
func isMatchV1(s string, p string) bool {
    switch {
    case len(p) == 0:
        return len(s) == 0
    case len(s) == 0:
        if len(p) > 1 && p[1] == '*' {
            return isMatchV1(s, p[2:])
        }
        return false
    case len(p) > 1 && p[1] == '*':
        if isMatchV1(s, p[2:]) {
            return true
        } else if s[0] == p[0] || p[0] == '.' {
            return isMatchV1(s[1:], p)
        }
        return false
    default:
        return (s[0] == p[0] || p[0] == '.') && isMatchV1(s[1:], p[1:])
    }
}
```