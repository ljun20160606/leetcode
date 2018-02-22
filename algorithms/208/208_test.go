package algorithms

import (
	"fmt"
	"testing"
)

func Test208(t *testing.T) {
	obj := Constructor()
	obj.Insert("aaa")
	param2 := obj.Search("aaa")
	param3 := obj.StartsWith("a")
	fmt.Println(param2, param3)
}
