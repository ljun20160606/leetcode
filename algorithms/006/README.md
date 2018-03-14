# ZigZag Conversion

## 描述

```txt

The string &#34;PAYPALISHIRING&#34; is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
P   A   H   N
A P L S I I G
Y   I   R


And then read line by line: &#34;PAHNAPLSIIGYIR&#34;


Write the code that will take a string and make this conversion given a number of rows:

string convert(string text, int nRows);

convert(&#34;PAYPALISHIRING&#34;, 3) should return &#34;PAHNAPLSIIGYIR&#34;.

```

## 题解

```go
package algorithms

//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//And then read line by line: "PAHNAPLSIIGYIR"
//Write the code that will take a string and make this conversion given a number of rows:
//
//string convert(string text, int nRows);
//convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

// 0123456789
// rows = 3
// 0   4   8
// 1 3 5 7 9
// 2   6

// rows = 4
// 0     6
// 1   5 7
// 2 4   8
// 3     9

// 不看规律纯插入字符
// direct = 1 代表向下逐个插入 direct = -1 代表向上逐个插入
// row 代表插哪一行 row = 0 direct = 1 row = rows - 1 direct = -1
func convert(s string, numRows int) string {
	if numRows < 2 || len(s) <= numRows {
		return s
	}
	rows := make([][]byte, numRows)
	for i := range rows {
		rows[i] = []byte{}
	}
	var direct, row int
	bottom := numRows - 1
	for _, v := range s {
		rows[row] = append(rows[row], byte(v))
		if row == 0 {
			direct = 1
		} else if row == bottom {
			direct = -1
		}
		row += direct
	}
	result := make([]byte, 0, len(s))
	for _, v := range rows {
		result = append(result, v...)
	}
	return string(result)
}

```