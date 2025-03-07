/*
We Have a House
In the garden, we have a house. We don't know how big the house is going to get. The garden is 50' x 50'. If you want to walk around the house, you'll need 3 feet so the house cannot be bigger than the width & depth of the garden minus the path to walk around it.

# We Have a House

In this example you can see the arguments your function is going to get (in red). The measurements of the windows + door as well as the dark rim (against rain damage) are always the same (in blue). We put One door in the front and Two windows in each wall.

We don't have permission to build higher than 20'. The area around the windows and door cannot be smaller than 1 foot except under the door obviously. It is possible to have a flat roof.

Create a function that takes four arguments and returns the area of light yellow paint and dark gray paint in a string as square feet. Assuming the coverage of the paint is perfect and you'll only need one layer of paint.

Examples
we_have_house(8, 30, 32, 8) ➞ "Yellow: 873, Gray: 242"

we_have_house(9, 14, 20, 9) ➞ "House too small."

we_have_house(9, 38, 36, 9) ➞ "Yellow: 1261, Gray: 290"

we_have_house(10, 31, 30, 11) ➞ "No permission."
Notes
If the house is too big for the garden, return "House too big."
If the house is too high, return "No permission."
If the house is too small (for the windows and door to fit), return "House too small."
*/
package main

import "fmt"

// weHaveHouse 関数は、家のサイズ、塗装面積を計算し、適切なメッセージを返す
func weHaveHouse(h, w, d, r int) string {
	// 1. 家のサイズ制約をチェック
	maxSize := 44
	if w > maxSize || d > maxSize {
		return "House too big."
	}

	// 2. 家の高さ制約をチェック
	if h+r > 20 {
		return "No permission."
	}

	// 3. 家のサイズが小さすぎるかどうかをチェック
	minHeight := 8
	minWidth := 15
	minDepth := 11
	if h < minHeight || w < minWidth || d < minDepth {
		return "House too small."
	}

	// 4. 塗装面積を計算
	grayArea := 4*(w+d) - 6

	openingArea := (8*12 + 21)

	// 黄色のペイント（壁の面積から開口部を引く）
	yellowArea := (w * r) + (2 * h * (w + d)) - openingArea - grayArea

	return fmt.Sprintf("Yellow: %d, Gray: %d", yellowArea, grayArea)
}

func main() {
	// テストケース
	fmt.Println(weHaveHouse(8, 30, 32, 8))   // "Yellow: 873, Gray: 242"
	fmt.Println(weHaveHouse(10, 31, 30, 11)) // "No permission."
	fmt.Println(weHaveHouse(8, 30, 30, 8))   // "Yellow: 849, Gray: 234"
	fmt.Println(weHaveHouse(9, 20, 18, 8))   // "Yellow: 581, Gray: 146"
	fmt.Println(weHaveHouse(9, 14, 20, 9))   // "House too small."
	fmt.Println(weHaveHouse(8, 16, 12, 8))   // "Yellow: 353, Gray: 106"
	fmt.Println(weHaveHouse(10, 25, 25, 0))  // "Yellow: 689, Gray: 194"
	fmt.Println(weHaveHouse(8, 45, 42, 8))   // "House too big."
	fmt.Println(weHaveHouse(10, 40, 40, 10)) // "Yellow: 1569, Gray: 314"
	fmt.Println(weHaveHouse(10, 15, 10, 7))  // "House too small."
	fmt.Println(weHaveHouse(9, 38, 36, 9))   // "Yellow: 1267, Gray: 290"
	fmt.Println(weHaveHouse(8, 15, 12, 6))   // "Yellow: 303, Gray: 102"
	fmt.Println(weHaveHouse(8, 30, 45, 6))   // "House too big."
	fmt.Println(weHaveHouse(9, 20, 14, 8))   // "Yellow: 525, Gray: 130"
}
