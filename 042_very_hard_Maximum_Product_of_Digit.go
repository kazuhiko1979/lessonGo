/*
Maximum Product of Digits
Write a function that returns all numbers less than or equal to N with the maximum product of digits.

Examples
max_product(8) ➞ [8]

max_product(27) ➞ [27]

max_product(211) ➞ [99, 199]

max_product(9578) ➞ [8999]
Notes
Search for numbers in the range: [0, n].
*/
package main

import (
	"fmt"
	"strconv"
)

func maxProduct(n int) []int {
	var maxProductNumbers []int
	maxProductValue := 0

	for i := range n + 1 {
		product := 1
		numStr := strconv.Itoa(i)

		for _, char := range numStr {
			digit, _ := strconv.Atoi(string(char))
			product *= digit
		}

		if product > maxProductValue {
			maxProductValue = product
			maxProductNumbers = []int{i}
		} else if product == maxProductValue {
			maxProductNumbers = append(maxProductNumbers, i)
		}
	}
	return maxProductNumbers
}

func main() {
	testCases := []int{8, 21, 26, 27, 43, 90, 199, 211, 455, 917, 1578, 9578, 11111, 41111}
	for _, test := range testCases {
		fmt.Println(maxProduct(test))
	}

}
