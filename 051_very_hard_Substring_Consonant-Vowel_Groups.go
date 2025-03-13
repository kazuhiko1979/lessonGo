/*
Substring Consonant-Vowel Groups
Write two functions:

One to retrieve all unique substrings that start and end with a vowel.
One to retrieve all unique substrings that start and end with a consonant.
The resulting array should be sorted in lexicographic ascending order (same order as a dictionary).

Examples
get_vowel_substrings("apple")
➞ ["a", "apple", "e"]

get_vowel_substrings("hmm") ➞ []

get_consonant_substrings("aviation")
➞ ["n", "t", "tion", "v", "viat", "viation"]

get_consonant_substrings("motor")
➞ ["m", "mot", "motor", "r", "t", "tor"]
Notes
Remember the output array should have unique values.
The word itself counts as a potential substring.
Exclude the empty string when outputting the array.
*/

package main

import (
	"fmt"
	"sort"
)

func getSubstrings(s string, isVowel bool) []string {
	vowels := "aeiou"
	substrings := make(map[string]struct{})

	for i := 0; i < len(s); i++ {

		if contains(vowels, s[i]) == isVowel {
			for j := i; j < len(s); j++ {
				if contains(vowels, s[j]) == isVowel {
					substrings[s[i:j+1]] = struct{}{}
				}
			}
		}
	}

	result := make([]string, 0, len(substrings))
	for substring := range substrings {
		result = append(result, substring)
	}
	sort.Strings(result)
	return result
}

func getVowelSubstrings(s string) []string {
	return getSubstrings(s, true)
}

func getConsonantSubstrings(s string) []string {
	return getSubstrings(s, false)
}

func contains(str string, char byte) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(getVowelSubstrings("apple"))        // ["a", "apple", "e"]
	fmt.Println(getVowelSubstrings("hmm"))          // []
	fmt.Println(getConsonantSubstrings("aviation")) // ["n", "t", "tion", "v", "viat", "viation"]
	fmt.Println(getConsonantSubstrings("motor"))    // ["m", "mot", "motor", "r", "t", "tor"]
}
