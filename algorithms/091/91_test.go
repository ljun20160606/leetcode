package algorithms

import "testing"

func Test91(t *testing.T) {
	decodings := numDecodings("1345")
	if decodings != 2 {
		t.Fatal("decodings should be 2")
	}
}
