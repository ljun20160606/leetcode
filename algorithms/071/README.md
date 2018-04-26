# Simplify Path

## 描述

```txt
Given an absolute path for a file (Unix-style), simplify it.

For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"

click to show corner cases.

Corner Cases:

 

 


	Did you consider the case where path = "/../"?
	In this case, you should return "/".
	Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
	In this case, you should ignore redundant slashes and return "/home/foo".


```

## 题解

```go
package algorithms

import "strings"

func simplifyPath(path string) string {
	var stack []string
	ss := strings.Split(path, "/")
	for i := range ss {
		s := ss[i]
		switch s {
		case "", ".":
		case "..":
			if len(stack) == 0 {
				continue
			}
			end := len(stack) - 1
			stack = stack[:end]
		default:
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}

```