package main

import "fmt"

// easterDate calculates the date of Easter for a given year (year >= 1600)
// using Butcher's Algorithm.

func easterDate(year int) string {
	// ステップ 1: 19年周期の月のサイクルを確認する
	// 年を19で割った余りを求める（月の満ち欠けの位置）
	a := year % 19

	// ステップ 2: 世紀と閏年の調整を行う
	// 年を100で割り、その商と余りを使い、閏年や世紀ごとの補正を計算する
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3

	// ステップ 3: パスカル満月のタイミングを計算する
	// これまでの結果を用いて、3月21日以降に最初に現れる満月の日を求める
	h := (19*a + b - d - g + 15) % 30

	// ステップ 4: 日曜日に合わせる調整
	// イースターは必ず日曜日であるため、計算した日付がどの曜日にあたるかを判定する
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7

	// ステップ 5: 最終的な計算で月と日の決定
	// 計算結果を基に、最終的にイースターが何月何日かを求める
	m := (a + 11*h + 22*l) / 451
	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	// 結果を文字列として返す
	var monthName string
	if month == 3 {
		monthName = "March"
	} else {
		monthName = "April"
	}
	return fmt.Sprintf("%s %d", monthName, day)
}

// func main() {

// 	fmt.Println(easterDate(1600)) // ➞ "April 2"
// 	fmt.Println(easterDate(2020)) // ➞ "April 12"
// 	fmt.Println(easterDate(1853)) // ➞ "March 27"
// 	fmt.Println(easterDate(3535)) // ➞ "April 14"

// }
