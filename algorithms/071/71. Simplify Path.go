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
