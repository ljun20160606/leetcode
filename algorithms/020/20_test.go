package algorithms

import "testing"

func Test20(t *testing.T) {
	parenthesis := "([{}])"
	valid := isValid(parenthesis)
	if !valid {
		t.Fatal("must be valid")
	}
	validV0 := isValidV0(parenthesis)
	if !validV0 {
		t.Fatal("must be valid")
	}
}
