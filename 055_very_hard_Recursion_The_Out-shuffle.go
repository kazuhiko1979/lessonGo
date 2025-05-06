/*
Recursion: The Out-Shuffle
An out-shuffle, also known as an out faro shuffle or a perfect shuffle, is a controlled method for shuffling playing cards. It is performed by splitting the deck into two equal halves and interleaving them together perfectly, with the condition that the top card of the deck remains in place.

Using an array to represent a deck of cards, an out-shuffle looks like:

[1, 2, 3, 4, 5, 6, 7, 8] ➞ [1, 5, 2, 6, 3, 7, 4, 8]
// Card 1 remains in the first position.
If we repeat the process, the deck eventually returns to original order:

Shuffle 1:

[1, 2, 3, 4, 5, 6, 7, 8] ➞ [1, 5, 2, 6, 3, 7, 4, 8]
Shuffle 2:

[1, 5, 2, 6, 3, 7, 4, 8] ➞ [1, 3, 5, 7, 2, 4, 6, 8]
Shuffle 3:

[1, 3, 5, 7, 2, 4, 6, 8] ➞ [1, 2, 3, 4, 5, 6, 7, 8]
// Back where we started.
Write a function that takes a positive even integer representing the number of the cards in a deck, and returns the number of out-shuffles required to return the deck to its original order.

Examples
shuffle_count(8) ➞ 3

shuffle_count(14) ➞ 12

shuffle_count(52) ➞ 8
Notes
The number of cards is always even and greater than one. Thus, the smallest possible deck size is two.
An iterative version of this challenge can be found here.
*/

package main

import (
	"reflect"
)

// アウトシャッフルを何回したら元に戻るかを調べる関数
func shuffleCount(n int) int {
	// 最初のカードの並びを作る（例：n=8なら [1, 2, 3, ..., 8]）
	original := make([]int, n)
	for i := 0; i < n; i++ {
		original[i] = i + 1
	}

	// 今のデッキの状態（最初は original と同じ）
	deck := make([]int, n)
	copy(deck, original)

	count := 0 // シャッフル回数

	for {
		count++

		half := n / 2
		left := deck[:half]  // 左半分
		right := deck[half:] // 右半分

		shuffled := make([]int, 0, n)
		for i := 0; i < half; i++ {
			shuffled = append(shuffled, left[i], right[i])
		} // シャッフル後のデッキ

		deck = shuffled // デッキを更新

		if reflect.DeepEqual(deck, original) {
			break
		}
	}
	return count
}

// func main() {
// 	fmt.Println(shuffleCount(8))  // ➞ 3
// 	fmt.Println(shuffleCount(14)) // ➞ 12
// 	fmt.Println(shuffleCount(52)) // ➞ 8
// }
