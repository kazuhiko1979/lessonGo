/*
Symmetrical Patterns
Kathleen owns a beautiful rug store. She likes to group the rugs into 4 mutually exclusive categories.

imperfect
horizontally symmetric
vertically symmetric
perfect
An imperfect rug is one that is neither horizontally nor vertically symmetric. Here is an example of an imperfect rug:

[
  ["a", "a", "a", "a"],
  ["a", "a", "a", "a"],
  ["a", "a", "b", "b"]
]
The following is an horizontally symmetric rug. You could "fold" the rug across a hypothetical x-axis, and both sides would be identical. A horizontally symmetric rug is not vertically symmetric (otherwise this rug would be classified as perfect ).

[
  ["c", "a", "a", "a"],
  ["b", "b", "b", "b"],
  ["c", "a", "a", "a"]
]
The following is a vertically symmetric rug. You could "fold" the rug across a hypothetical y-axis, and both sides would be identical. A vertically symmetric is not horizontally symmetric (otherwise this rug would be classified as perfect ).

[
  ["a", "b", "a"],
  ["b", "b", "b"],
  ["a", "b", "a"],
  ["a", "b", "a"]
]
Finally, a perfect rug is one that is both vertically and horizontally symmetric. That is, folded either length-wise or width-wise will yield two identical pieces.

[
  ["a", "b", "b", "a"],
  ["b", "b", "b", "b"],
  ["a", "b", "b", "a"]
]
Given a rug of m x n dimension, determine whether it is imperfect, horizontally symmetric, vertically symmetric or perfect. Rugs are represented using a two-dimensional list.

Examples
classify_rug([
  ["a", "a"],
  ["a", "a"]
]) ➞ "perfect"

classify_rug([
  ["a", "a", "b"],
  ["a", "a", "a"],
  ["b", "a", "a"]
]) ➞ "imperfect"

classify_rug([
  ["b", "a"],
  ["b", "a"]
]) ➞ "horizontally symmetric"

classify_rug([
  ["a", "a"],
  ["b", "b"]
]) ➞ "vertically symmetric"
Notes
You can consider a 1 x n rug as being trivially horizontally symmetric, an n x 1 rug as being trivially vertically symmetric, and a 1 x 1 rug as being trivially perfect.
*/

package main

import "fmt"

// 横に対象かどうか
func isHorizontallySymmetric(rug [][]string) bool {
	m := len(rug)
	for i := 0; i < m/2; i++ {
		if fmt.Sprint(rug[i]) != fmt.Sprint(rug[m-1-i]) {
			return false
		}
	}
	return true
}

// 縦に対象かどうか

func isVerticallySymmertric(rug [][]string) bool {
	m, n := len(rug), len(rug[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			if rug[i][j] != rug[i][n-1-j] {
				return false
			}
		}
	}
	return true
}

// ラグを分類
func classifyRug(rug [][]string) string {
	horizontal := isHorizontallySymmetric(rug)
	vertical := isVerticallySymmertric(rug)

	if horizontal && vertical {
		return "perfect"
	} else if horizontal {
		return "horizontally symmetric"
	} else if vertical {
		return "vertically symmetric"
	} else {
		return "imperfect"
	}
}

// func main() {
// 	// テストケース
// 	fmt.Println(classifyRug([][]string{
// 		{"a", "a"},
// 		{"a", "a"},
// 	})) // ➞ "perfect"

// 	fmt.Println(classifyRug([][]string{
// 		{"a", "a", "b"},
// 		{"a", "a", "a"},
// 		{"b", "a", "a"},
// 	})) // ➞ "imperfect"

// 	fmt.Println(classifyRug([][]string{
// 		{"b", "a"},
// 		{"b", "a"},
// 	})) // ➞ "horizontally symmetric"

// 	fmt.Println(classifyRug([][]string{
// 		{"a", "a"},
// 		{"b", "b"},
// 	})) // ➞ "vertically symmetric"
// }
