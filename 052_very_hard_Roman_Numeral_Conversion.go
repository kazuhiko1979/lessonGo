/*
Roman Numeral Conversion
Create a function that will take either a string containing a roman numeral, or an integer.

Given a string, return the integer value of that roman numeral.
Given an integer, return the equivalent roman numeral.
Symbols to Values
I ➞ 1

V ➞ 5

X ➞ 10

L ➞ 50

C ➞ 100

D ➞ 500

M ➞ 1000
Examples
roman_numerals("V") ➞ 5

roman_numerals("IV") ➞ 4

roman_numerals("CII") ➞ 102

roman_numerals(45) ➞ "XLV"

roman_numerals(1666) ➞ "MDCLXVI"
Notes
Numerical and Roman Numeral inputs will be positive numbers between 1 and 9999.
*/

package main

import (
	"fmt"
	"strings"
)

func romanNumerals(input interface{}) interface{} {
	romanDict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}

	valueMap := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	switch v := input.(type) {
	case string:
		total := 0
		prevValue := 0
		for i := len(v) - 1; i >= 0; i-- {
			currentValue := romanDict[string(v[i])]
			if currentValue < prevValue {
				total -= currentValue
			} else {
				total += currentValue
			}
			prevValue = currentValue
		}
		return total

	case int:
		var romanStr strings.Builder
		for _, pair := range valueMap {
			for v >= pair.Value {
				romanStr.WriteString(pair.Symbol)
				v -= pair.Value
			}
		}
		return romanStr.String()
	}
	return nil
}

func main() {
	testCases := []interface{}{
		"I", "V", "X", "L", "C", "D", "M", "IV", "VI", "XIV", "LIX", "XCIX",
		"CII", "XLV", "XXX", "XXXVI", "DCCXIV", "MMXVIII", "MDCLXVI", "MCCCXXIV", 1324,
	}

	for _, testCase := range testCases {
		fmt.Println(romanNumerals(testCase))
	}
}
