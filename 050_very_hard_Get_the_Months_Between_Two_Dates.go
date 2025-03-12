/*
Get the Months Between Two Dates
Create a function that, given 2 dates, returns the names of the months that are present between them (inclusive).

Examples
Input

 january = datetime.date(2017, 1, 1)
 march = datetime.date(2017, 3, 1)

monthsInterval(january, march)
Output

['January', 'February', 'March']
Input

 december = datetime.date(2017, 12, 1)
 january = datetime.date(2018, 1, 1)

monthsInterval(december, january)
Output

['January', 'December']
Input

 january2017 = datetime.date(2017, 0, 1)
 january2018 = datetime.date(2018, 0, 1)

monthsInterval(january2017, january2018)
Output

['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
(Notice that January is not duplicated!)

Notes
The returned list should include the months of dateStart and dateEnd (inclusive)
The returned list must not include duplicate values
The month names returned by the function should be sorted (not alphabetically, but ordered by which comes first (January = 1st month, February = 2nd month, … , December = 12th month))
The function should produce the same output even if dateStart is greater than dateEnd
*/

package main

import (
	"fmt"
	"sort"
	"time"
)

// monthsIntervalは、2つの日付の間に登場する月の名前を返す関数です。
// 日付の順序が逆の場合は、入れ替えます。

func monthsInterval(date1, date2 time.Time) []string {
	// 1. 受け取った2つの日付のうち、早い方をstart、遅い方をendにする
	if date1.After(date2) {
		date1, date2 = date2, date1
	}

	// 2. 1月から12月までの月の名前の一覧を作成する
	monthNames := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	// 結果を入れるスライス
	var result []string

	monthExists := make(map[string]bool)

	// 3. startの日付からendの日付まで、月ごとにループするための準備
	currentYear := date1.Year()
	currentMonth := int(date1.Month())

	// ループ条件：(currentYear, currentMonth) が (date2.Year(), date2.Month()) 以下である間ループ
	for currentYear < date2.Year() || (currentYear == date2.Year() && currentMonth <= int(date2.Month())) {
		// 4. 現在の月の名前を取得し、重複がなければ結果に追加する
		month := monthNames[currentMonth-1]
		if !monthExists[month] {
			result = append(result, month)
			monthExists[month] = true
		}

		// 5. 次の月に進む（12月の場合は翌年の1月へ）
		if currentMonth == 12 {
			currentMonth = 1
			currentYear++
		} else {
			currentMonth++
		}
	}

	// 6. 結果のスライスをカレンダー順（1月～12月）に並び替えるため、monthNames内のインデックスを基準にソートする
	sort.Slice(result, func(i, j int) bool {
		return indexOf(monthNames, result[i]) < indexOf(monthNames, result[j])
	})

	return result
}

// indexOfは、スライス内から文字列sのインデックスを返す関数です。
// 見つからなければ-1を返します。
func indexOf(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}
	return -1
}

func main() {
	// テスト例1
	january := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	march := time.Date(2017, time.March, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Test 1:", monthsInterval(january, march))

	// テスト例2
	december := time.Date(2017, time.December, 1, 0, 0, 0, 0, time.UTC)
	januaryNext := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Test 2:", monthsInterval(december, januaryNext))

	// テスト例3
	january2017 := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	january2018 := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Test 3:", monthsInterval(january2017, january2018))
}
