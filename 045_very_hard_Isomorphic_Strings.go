package main

import (
	"fmt"
)

// isIsomorphic は、文字列 s と t がイソモルフィック（
// 「sの文字 ⇔ tの文字」が 1 対 1 で矛盾なく対応するか）を判定します。
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 2. rune スライスに変換して、rune単位でも長さを比較
	sr := []rune(s)
	tr := []rune(t)
	if len(sr) != len(tr) {
		return false
	}

	// 3. 文字列を1文字ずつ見ていき、対応表を作る
	// 例: s="egg", t="add" の場合、	{'e': 'a', 'g': 'd'}
	// 例: s="foo", t="bar" の場合、	{'f': 'b', 'o': 'a'}
	// 例: s="paper", t="title" の場合、{'p': 't', 'a': 'i', 'e': 'l', 'r': 'e'}
	sToT := make(map[rune]rune)
	tToS := make(map[rune]rune)

	// 4. sr と tr を同時に走査し、対応関係を検証
	for i, rS := range sr {
		rT := tr[i]

		if mapped, ok := sToT[rS]; ok {
			if mapped != rT {
				return false
			}
		} else {
			// 新しく対応関係を登録する際、rT がほかの文字に使われていないか確認
			if rev, exists := tToS[rT]; exists && rev != rS {
				return false
			}
			sToT[rS] = rT
			tToS[rT] = rS
		}
	}
	return true
}

// 確認用
func main() {
	fmt.Println(isIsomorphic("egg", "add"))                 // => true
	fmt.Println(isIsomorphic("aba", "baa"))                 // => false
	fmt.Println(isIsomorphic("paper", "title"))             // => true
	fmt.Println(isIsomorphic("ab", "ca"))                   // => true
	fmt.Println(isIsomorphic("abab", "baba"))               // => true
	fmt.Println(isIsomorphic("approximate", "tooxaezptom")) // => false
}
