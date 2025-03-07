/*
Hexadecimal Color Mixer
In this challenge, you have to mix two or more colors to get a brand new color from their average rgb values.

Each color will be represented in its hexadecimal notation, and so as a string starting with # containing three pairs of alphanumeric characters, equals to the three rgb values (in base 16) of red, green and blue.

To obtain the new color, you need to get the arithmetic average of the sums of the individual red, green and blue values of each given color. When the average is a float number, you have to round it to the nearest integer (rounding up for decimal parts equal to 0.5).

Mixing yellow and red:

Hexadecimal code of yellow = "#FFFF00"
Hexadecimal code of red = "#FF0000"

yellow to RGB = Red: ["FF" = 255], Green: ["FF" = 255], Blue: ["00", 0]
red to RGB = Red: ["FF" = 255], Green: ["00" = 0], Blue: ["00" = 0]

Average of Red values = (255 + 255) / 2 = 255
Average of Green values = (255 + 0) / 2 = 127.5 = 128
Average of Blue values = (0 + 0) / 2 = 0 = 0

Rgb of new color = [255, 128, 0]
Hexadecimal conversion = [255 = "FF"], [128 = "80"], [0 = "00"]

New color = "#FF8000" (orange)
Given a list of strings colors containing a series of hexadecimal codes, implement a function that returns the hexadecimal code of the new color, as a new string.

Examples
hex_color_mixer(["#FFFF00", "#FF0000"]) ➞ "#FF8000"
# Orange

hex_color_mixer(["#FFFF00", "#0000FF"]) ➞ "#808080"
# Medium gray

hex_color_mixer(["#B60E73", "#0EAEB6"]) ➞ "#625E95"
# Lavender
Notes
Remember to round to the nearest integer the average of the rgb values.
You can test the color codes in any hexadecimal-colors utility site, such as this one that I used for testing cases.
All the given hexadecimal strings are valid; there are no exceptions to handle.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// hexColorMixer は #RRGGBB 形式の色コードを複数受け取り、
// 平均値を #RRGGBB 形式で返す。

func hexColorMixer(colors []string) string {
	var totalR, totalG, totalB float64

	// 1. それぞれの色から R, G, B の値を取り出して合計
	for _, color := range colors {
		color = strings.TrimPrefix(color, "#")
		r, _ := strconv.ParseInt(color[0:2], 16, 64)
		g, _ := strconv.ParseInt(color[2:4], 16, 64)
		b, _ := strconv.ParseInt(color[4:6], 16, 64)

		totalR += float64(r)
		totalG += float64(g)
		totalB += float64(b)
	}

	// 2. 各要素の平均値を出して、小数点以下がちょうど 0.5 なら切り上げる
	count := float64(len(colors))
	avgR := int(math.Floor(totalR/count + 0.5))
	avgG := int(math.Floor(totalG/count + 0.5))
	avgB := int(math.Floor(totalB/count + 0.5))

	return fmt.Sprintf("#%02X%02X%02X", avgR, avgG, avgB)
}

// func main() {
// 	fmt.Println(hexColorMixer([]string{"#FFFF00", "#FF0000"}))
// 	fmt.Println(hexColorMixer([]string{"#FFFF00", "#0000FF"}))                                             // "#808080", Example #2
// 	fmt.Println(hexColorMixer([]string{"#B60E73", "#0EAEB6"}))                                             // "#625E95", Example #3
// 	fmt.Println(hexColorMixer([]string{"#FF0000", "#00FF00", "#0000FF"}))                                  // "#555555"
// 	fmt.Println(hexColorMixer([]string{"#99CC00", "#663399", "#1A8191"}))                                  // "#5E8063"
// 	fmt.Println(hexColorMixer([]string{"#918381", "#D53B21", "#A54C83", "#DEFACF"}))                       // "#BA817D"
// 	fmt.Println(hexColorMixer([]string{"#140A23", "#46B31E", "#CFDF08", "#263534", "#EAD1FB", "#235E02"})) // "#65803F"
// 	fmt.Println(hexColorMixer([]string{"#FFFFFF", "#000000", "#000000", "#FFFFFF"}))                       // "#808080"
// 	fmt.Println(hexColorMixer([]string{"#CCCCCC"}))
// }
