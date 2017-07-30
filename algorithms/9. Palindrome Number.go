package algorithms

//Determine whether an integer is a palindrome. Do this without extra space.

// 反转数字 判断是否相等
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	t := x
	var r int
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return t == r
}
