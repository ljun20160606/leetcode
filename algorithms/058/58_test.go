package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	ast := assert.New(t)

	qst := struct {
		input  string
		expect int
	}{
		input:  "Hello World",
		expect: 5,
	}

	ast.Equal(qst.expect, lengthOfLastWord(qst.input), "输入:%v", qst.input)
}
