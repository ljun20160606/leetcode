package algorithms

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && ignore(s[left]) {
			left++
		}
		for left < right && ignore(s[right]) {
			right--
		}

		if lowerCase(s[left]) != lowerCase(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func ignore(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return false
	}
	return true
}

func lowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 'a' - 'A'
	}
	return b
}
