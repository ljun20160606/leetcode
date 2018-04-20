package algorithms

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	if version1 == version2 {
		return 0
	}
	p1, version1 := con(version1)
	p2, version2 := con(version2)
	if p1 > p2 {
		return 1
	} else if p1 < p2 {
		return -1
	}
	return compareVersion(version1, version2)
}

func con(version string) (number int, newVersion string) {
	sp := strings.SplitN(version, ".", 2)
	if len(sp) == 1 {
		p, _ := strconv.Atoi(sp[0])
		return p, ""
	}
	p, _ := strconv.Atoi(sp[0])
	return p, sp[1]
}
