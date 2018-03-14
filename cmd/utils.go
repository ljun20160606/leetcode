package main

import "strings"

// 填补 e.g. 1 -> 001
func offsetNumber(c string, offsetByte byte, amount int) string {
	if len(c) > amount {
		return c
	}
	missing := amount - len(c)
	builder := strings.Builder{}
	for i := 0; i < missing; i++ {
		builder.Write([]byte{offsetByte})
	}
	builder.WriteString(c)
	return builder.String()
}
