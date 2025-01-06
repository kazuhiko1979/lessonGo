package main

import (
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

// func main() {
// 	// テスト用の出力
// 	a, b := simplifySqrt(72)
// 	fmt.Printf("(%d, %d)\n", a, b) // ➞ (6, 2)

// 	a, b = simplifySqrt(160)
// 	fmt.Printf("(%d, %d)\n", a, b) // ➞ (4, 10)

// 	a, b = simplifySqrt(36)
// 	fmt.Printf("(%d, %d)\n", a, b) // ➞ (6, 1)

// 	a, b = simplifySqrt(35)
// 	fmt.Printf("(%d, %d)\n", a, b) // ➞ (1, 35)
// }
