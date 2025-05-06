/*
Find the Pattern and Write the Function
By looking at the inputs and outputs below, try to figure out the pattern and write a function to execute it for any number.

Examples
func(3456) ➞ 2

func(89265) ➞ 5

func(97) ➞ 12

func(2113) ➞ -9"

*/

package main

import (
	"fmt"
	"strconv"
)

func funcDigits(n int) int {
	s := strconv.Itoa(n)

	sum := 0
	for _, ch := range s {
		sum += int(ch - '0')
	}

	k := len(s)
	return sum - k*k
}
func main() {
	// テスト出力
	fmt.Println(funcDigits(3456))   // 2
	fmt.Println(funcDigits(89265))  // 5
	fmt.Println(funcDigits(97))     // 12
	fmt.Println(funcDigits(2113))   // -9
	fmt.Println(funcDigits(55))     // 6
	fmt.Println(funcDigits(785428)) // -2
	fmt.Println(funcDigits(439))    // 7
	fmt.Println(funcDigits(55654))  // 0
}
