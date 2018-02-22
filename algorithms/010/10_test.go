package algorithms

import "testing"

func Test10(t *testing.T) {
	if isMatch("aa", "a") {
		t.Fatal("should be false")
	}
	if !isMatch("aa", "aa") {
		t.Fatal("should be true")
	}
	if isMatch("aaa", "aa") {
		t.Fatal("should be false")
	}
	if !isMatch("aa", "a*") {
		t.Fatal("should be true")
	}
	if !isMatch("aa", ".*") {
		t.Fatal("should be true")
	}
	if !isMatch("ab", ".*") {
		t.Fatal("should be true")
	}
	if !isMatch("aab", "c*a*b") {
		t.Fatal("should be true")
	}
}

func Test10V1(t *testing.T) {
	if isMatchV1("aa", "a") {
		t.Fatal("should be false")
	}
	if !isMatchV1("aa", "aa") {
		t.Fatal("should be true")
	}
	if isMatchV1("aaa", "aa") {
		t.Fatal("should be false")
	}
	if !isMatchV1("aa", "a*") {
		t.Fatal("should be true")
	}
	if !isMatchV1("aa", ".*") {
		t.Fatal("should be true")
	}
	if !isMatchV1("ab", ".*") {
		t.Fatal("should be true")
	}
	if !isMatchV1("aab", "c*a*b") {
		t.Fatal("should be true")
	}
}

func Test10V0(t *testing.T) {
	if isMatchV0("aa", "a") {
		t.Fatal("should be false")
	}
	if !isMatchV0("aa", "aa") {
		t.Fatal("should be true")
	}
	if isMatchV0("aaa", "aa") {
		t.Fatal("should be false")
	}
	if !isMatchV0("aa", "a*") {
		t.Fatal("should be true")
	}
	if !isMatchV0("aa", ".*") {
		t.Fatal("should be true")
	}
	if !isMatchV0("ab", ".*") {
		t.Fatal("should be true")
	}
	if !isMatchV0("aab", "c*a*b") {
		t.Fatal("should be true")
	}
}
