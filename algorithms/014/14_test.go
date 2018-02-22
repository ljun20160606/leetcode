package algorithms

import "testing"

func Test14(t *testing.T) {
	a := longestCommonPrefix([]string{"ab", "a", "ac", "abd"})
	if a != "a" {
		t.Fatal("result should be a")
	}
	none := longestCommonPrefix([]string{})
	if none != "" {
		t.Fatal("result should be ")
	}
}
