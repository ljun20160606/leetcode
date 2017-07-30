package algorithms

//Given a roman numeral, convert it to an integer.
//
//Input is guaranteed to be within the range from 1 to 3999.

// 罗马数字的规律
// VI 前一个数字大于后一个则相加
// IV 前一个数字小于后一个则相减
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := m[s[0]]
	for i := 1; i < len(s); i++ {
		if last, now := m[s[i-1]], m[s[i]]; last < now {
			result += now - 2*last
		} else {
			result += now
		}
	}
	return result
}
