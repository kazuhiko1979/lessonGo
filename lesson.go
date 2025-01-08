// Go:build main

package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	// // 最小値
	// min := l[0]
	// for _, v := range l {
	// 	if v < min {
	// 		min = v
	// 	}
	// }

	// fmt.Printf("最小値: %d\n", min)

	minValue := l[0]
	for _, v := range l {
		minValue = min(minValue, v)
	}

	fmt.Printf("最小値 %d\n", minValue)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grape":  150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	// 合計
	total := 0
	for _, price := range m {
		total += price
	}

	fmt.Println("果物の価格の合計: %d円\n", total)
}
