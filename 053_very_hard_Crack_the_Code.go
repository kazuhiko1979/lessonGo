/*
Crack the Code
This is a reverse-coding challenge. Create a function that outputs the correct list from the input. Use the following examples to crack the code.

Examples
decode("hello") ➞ [5, 2, 9, 9, 3]

decode("wonderful") ➞ [11, 3, 2, 1, 2, 6, 3, 9, 9]

*/

package main

import (
	"fmt"
	"strconv"
)

func decode(txt string) []int {
	var lst []int

	for _, letter := range txt {
		numOrd := int(letter) // Unicode (ASCII)
		sum := 0

		// 数値を文字列に変換し、各桁の和を計算
		strNum := strconv.Itoa(numOrd)
		for _, num := range strNum {
			digit, _ := strconv.Atoi(string(num))
			sum += digit
		}

		lst = append(lst, sum)
	}

	return lst
}

func main() {
	// テストケース
	fmt.Println(decode("hello"))          // [5, 2, 9, 9, 3]
	fmt.Println(decode("wonderful"))      // [11, 3, 2, 1, 2, 6, 3, 9, 9]
	fmt.Println(decode("all my friends")) // [16, 9, 9, 5, 10, 4, 5, 3, 6, 6, 2, 2, 1, 7]
	fmt.Println(decode("River"))          // [10, 6, 10, 2, 6]
}
