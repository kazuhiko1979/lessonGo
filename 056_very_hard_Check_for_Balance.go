/*
Check for Balance
Write a function that takes a string of source code and checks whether the braces/parentheses are Balanced. Every ( or { must be closed by a } or ) in the opposite order. Return the index at which an imBalance occurs, or -1 if the string is Balanced. If any ( or { are never closed, return the string's length.

Examples
check_Balance("if (a(4) > 9) { foo(a(2)); }") ➞ -1
# Returns -1 because it's Balanced.

check_Balance("for (i=0;i<a(3};i++) { foo{); )") ➞ 14
# Returns 14 because } is out of order.

check_Balance("if (x) {")  ➞ 8
# Returns 8 because { is never closed.
*/

package main

import "fmt"

func checkBalance(code string) int {
	type stackItem struct {
		char  rune
		index int
	}
	var stack []stackItem

	for i, ch := range code {
		switch ch {
		case '(', '{':
			stack = append(stack, stackItem{ch, i})
		case ')', '}':
			if len(stack) == 0 {
				return i
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (ch == ')' && top.char != '(') || (ch == '}' && top.char != '{') {
				return i
			}
		}
	}

	if len(stack) > 0 {
		return stack[0].index + 1
	}
	return -1
}

func main() {

	fmt.Println(checkBalance("if (a(4) > 9) { foo(a(2)); }"))
	fmt.Println(checkBalance("for (i=0;i<a(3};i++) { foo{); )"))
	fmt.Println(checkBalance("while (true) foo(); }{ ()"))
	fmt.Println(checkBalance("if (x) {"))
	fmt.Println(checkBalance("if (x) }"))
	fmt.Println(checkBalance("(({{}}){}{}())"))
	fmt.Println(checkBalance("({)}"))
	fmt.Println(checkBalance("("))
	fmt.Println(checkBalance("}"))
	fmt.Println(checkBalance(""))

}
