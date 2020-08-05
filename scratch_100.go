package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	length := len(s)
	if numRows <= 1 {
		return s
	}
	if length < numRows {
		numRows = length
	}

	rows := make([][]byte, numRows)
	var currRow int
	var isDown bool
	for i := 0; i < length; i++ {
		rows[currRow] = append(rows[currRow], s[i])
		if currRow == 0 || currRow == numRows-1 {
			isDown = !isDown
		}
		if isDown {
			currRow += 1
		} else {
			currRow -= 1
		}
	}
	result := make([]byte, 0, length)
	for i := 0; i < numRows; i++ {
		result = append(result, rows[i]...)
	}
	return string(result)
}

func main() {
	fmt.Println(convert("LEETCODEISHIRING", 0) == "LEETCODEISHIRING")
	fmt.Println(convert("LEETCODEISHIRING", 1) == "LEETCODEISHIRING")
	fmt.Println(convert("LEETCODEISHIRING", 3) == "LCIRETOESIIGEDHN")
	fmt.Println(convert("LEETCODEISHIRING", 4) == "LDREOEIIECIHNTSG")
}
