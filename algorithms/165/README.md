# Compare Version Numbers

## 描述

```txt
Compare two version numbers version1 and version2.
If version1 > version2 return 1, if version1 < version2 return -1, otherwise return 0.

You may assume that the version strings are non-empty and contain only digits and the . character.
The . character does not represent a decimal point and is used to separate number sequences.
For instance, 2.5 is not "two and a half" or "half way to version three", it is the fifth second-level revision of the second first-level revision.

Here is an example of version numbers ordering:

0.1 < 1.1 < 1.2 < 13.37

Credits:
Special thanks to @ts for adding this problem and creating all test cases.

```

## 题解

```go
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

```