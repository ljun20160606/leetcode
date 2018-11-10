package algorithms

import "strings"

func reverseString(s string) string {
	builder := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteByte(s[i])
	}
	return builder.String()
}
