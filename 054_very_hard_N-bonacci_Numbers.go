/*
N-bonacci numbers are generalisations of the fibonacci sequence, where the next term is always the sum of the previous N terms. By convention, the first (N-1) terms are all 0 and the Nth term is 1.

The initial 10 terms of the first 5 N-bonacci sequences are therefore:

1-bonacci = 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, ...
2-bonacci = 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
3-bonacci = 0, 0, 1, 1, 2, 4, 7, 13, 24, 44, ...
4-bonaaci = 0, 0, 0, 1, 1, 2, 4, 8, 15, 29, ...
5-bonacci = 0, 0, 0, 0, 1, 1, 2, 4, 8, 16, ...
Write a function that returns the kth term of the N-bonacci sequence, for two integer arguments N and k.

Examples
bonacci(1, 10) ➞ 1

bonacci(2, 10) ➞ 34

bonacci(3, 10) ➞ 44

bonacci(4, 10) ➞ 29

bonacci(5, 10) ➞ 16

*/

package main

func bonacci(N, k int) int {
	// kがNより前なら、初期値なので0を返す（最初のN-1項は0）
	if k <= N-1 {
		return 0
	}

	// 初期値：N-1個の0と、1を1つ
	sequence := make([]int, N-1)
	sequence = append(sequence, 1)

	// N番目以降は、直前のN個の和を追加していく
	for i := 0; i < k-N; i++ {
		sum := 0
		for j := len(sequence) - N; j < len(sequence); j++ {
			sum += sequence[j]
		}
		sequence = append(sequence, sum)
	}
	return sequence[k-1]
}

// func main() {
// 	fmt.Println(bonacci(2, 7))  // 8
// 	fmt.Println(bonacci(3, 13)) // 274
// 	fmt.Println(bonacci(5, 24)) // 203513
// 	fmt.Println(bonacci(8, 44)) // 32440904961（intオーバーフローに注意）
// 	fmt.Println(bonacci(1, 4))  // 1
// 	fmt.Println(bonacci(2, 2))  // 1
// 	fmt.Println(bonacci(3, 1))  // 0
// }
