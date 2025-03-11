/*
Make Anagram
Given two strings, that may or may not be of the same length, determine the minimum number of character deletions required to make an anagram. Any characters can be deleted from either of the strings.

Examples
make_anagram("cde", "abc") ➞ 4
# Remove d and e from cde to get c.
# Remove a and b from abc to get c.
# It takes 4 deletions to make both strings anagrams.

make_anagram("fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke") ➞ 30

make_anagram("showman", "woman") ➞ 2
*/

package main

import "fmt"

func makeAnagram(str1, str2 string) int {
	// 文字の出現回数をカウント
	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	for _, char := range str1 {
		count1[char]++
	}

	// str2の文字をカウント
	for _, char := range str2 {
		count2[char]++
	}

	// すべてのユニークな文字を取得
	uniqueChars := make(map[rune]bool)
	for char := range count1 {
		uniqueChars[char] = true
	}

	for char := range count2 {
		uniqueChars[char] = true
	}

	// 削除する文字数を計算
	deletions := 0
	for char := range uniqueChars {
		deletions += abs(count1[char] - count2[char])
	}

	return deletions
}

// 絶対値を求める
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	// テストケース
	fmt.Println(makeAnagram("cde", "abc"))                                             // ➞ 4
	fmt.Println(makeAnagram("fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke")) // ➞ 30
	fmt.Println(makeAnagram("showman", "woman"))                                       // ➞ 2
}
