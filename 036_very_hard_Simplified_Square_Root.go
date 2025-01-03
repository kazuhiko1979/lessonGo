/*
Simplified Square Root
Create a function that takes in a positive integer n and returns the simplified square root of n as a tuple of positive integers (a, b), where a⋅sqrt(b) = sqrt(n) and b is as small as possible.

Examples
simplify_sqrt(72) ➞ (6, 2)

simplify_sqrt(160) ➞ (4, 10)

simplify_sqrt(36) ➞ (6, 1)

simplify_sqrt(35) ➞ (1, 35)
A common way to simplify square roots is to repeatedly factor out perfect squares from the number underneath the square root. For example, if you need to simply sqrt(72), you can factor out perfect squares from 72 according to the following process:

sqrt(72)
72 is divisible by 4, so factor 4 out of 72:

sqrt(4⋅18)
The perfect square 4 can now be square rooted and pulled out of the square root:

2⋅sqrt(18)
Now repeat the process until no further perfect squares can be factored out. 18 is divisible by 9, so factor it out:

2⋅sqrt(9⋅2)
Pull the 9 out, square root it, and simplify:

2⋅3⋅sqrt(2)
= 6⋅sqrt(2)
2 does not have any perfect square factors other than 1, so 6⋅sqrt(2) is the simplest form of sqrt(72). The function would therefore return the tuple (6, 2).

This is only one approach to solving this problem; there are probably other ways that are simpler/faster than this method. Feel free to use any method you want.

Notes
If n is a perfect square, b should be 1.
If n has no perfect square factors (other than 1), a should be 1.
*/

package main

import (
	"fmt"
	"math"
)

func simplifySqrt(n int) (int, int) {
	maxSquare := 1

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		square := i * i
		if n%square == 0 {
			maxSquare = square
		}
	}

	a := int(math.Sqrt(float64(maxSquare)))
	b := n / maxSquare
	return a, b
}

func main() {
	// テスト用の出力
	a, b := simplifySqrt(72)
	fmt.Printf("(%d, %d)\n", a, b) // ➞ (6, 2)

	a, b = simplifySqrt(160)
	fmt.Printf("(%d, %d)\n", a, b) // ➞ (4, 10)

	a, b = simplifySqrt(36)
	fmt.Printf("(%d, %d)\n", a, b) // ➞ (6, 1)

	a, b = simplifySqrt(35)
	fmt.Printf("(%d, %d)\n", a, b) // ➞ (1, 35)
}
