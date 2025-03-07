/*
Ruth and Aaron
Two consecutive integers a and b are considered a Ruth-Aaron pair if the sum of the prime factors of a is equal to the sum of the prime factors of b.

Two definitions exist:

Summing up distinct prime factors. For example, 24 and 25 constitute a Ruth-Aaron pair by this definition. 8 and 9 do not.
P24 = [2, 3]  # sum = 5

P25 = [5]  # sum = 5, equal to 24

P8 = [2]  # sum = 2

P9 = [3]  # sum = 3
Summing up repeated prime factors. By this definition, 24 and 25 do not constitute a Ruth-Aaron pair. But 8 and 9 do.
P24 = [2, 2, 2, 3]  # sum = 9

P25 = [5, 5]  # sum = 10

P8 = [2, 2, 2]  # sum = 6

P9 = [3, 3]  # sum = 6, equal to 8
If two consecutive numbers have only distinct prime factors and have equal sums of prime factors, they are considered Ruth-Aaron pairs by both definitions.

P77 = [7, 11]  # sum = 18

P78 = [2, 3, 13]  # sum = 18
Create a function that takes a number n and returns:

False if it is not part of a Ruth-Aaron pair.
A 2-element list if it is part of a Ruth-Aaron pair.
The first element should be "Ruth" if n is the smaller number in the pair, or "Aaron" if it is the larger.
The second element should be 1 if n is part of a Ruth-Aaron pair under the first definition (sum of distinct prime factors)) # 2 if it qualifies by the second definition, 3 if it qualifies under both.
Examples
ruth_aaron(5) ➞ ["Ruth", 3]

ruth_aaron(25) ➞ ["Aaron", 1]

ruth_aaron(9) ➞ ["Aaron", 2]

ruth_aaron(11) ➞ False
*/

package main

// getPrimeFactors は、与えられた整数 x の素因数を重複を含めて列挙して返します。
// 例: 12 -> [2, 2, 3]

func getPrimeFactors(x int) []int {
	factors := []int{}
	num := x
	factor := 2

	for factor*factor <= num {
		for num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		}
		// 最初だけ 2 -> 3、それ以降は 2 ずつ増やして奇数を試す
		if factor == 2 {
			factor = 3
		} else {
			factor += 2
		}
	}
	if num > 1 {
		factors = append(factors, num)
	}
	return factors
}

// sumOfDistinctPrimeFactors は、与えられた整数 x の
// 「異なる素因数」の和を返します。
// 例: 24 -> [2, 3] (distinct) -> 和 = 5
func sumOfDistinctPrimeFactors(x int) int {
	factors := getPrimeFactors(x)
	distinct := make(map[int]bool)

	for _, f := range factors {
		distinct[f] = true
	}
	sum := 0
	for k := range distinct {
		sum += k
	}
	return sum
}

// sumOfPrimeFactorsWithMultiplicity は、与えられた整数 x の
// 「重複を含む素因数」の和を返します。
// 例: 8 -> [2, 2, 2] -> 和 = 6
func sumOfPrimeFactorsWithMultiplicity(x int) int {
	factors := getPrimeFactors(x)
	sum := 0
	for _, f := range factors {
		sum += f
	}
	return sum
}

// checkPair は、2 つの整数 a, b が連続しているとして、Ruth–Aaron の関係かを判定し
// 以下の整数値を返します。
//   0: ペアではない
//   1: 異なる素因数の和だけ一致
//   2: 重複ありの素因数の和だけ一致
//   3: 両方の定義で一致

func checkPair(a, b int) int {
	sdA := sumOfDistinctPrimeFactors(a)
	sdB := sumOfDistinctPrimeFactors(b)
	smA := sumOfPrimeFactorsWithMultiplicity(a)
	smB := sumOfPrimeFactorsWithMultiplicity(b)

	distinctMatch := (sdA == sdB)
	multiplicityMatch := (smA == smB)

	switch {
	case distinctMatch && multiplicityMatch:
		return 3
	case distinctMatch:
		return 1
	case multiplicityMatch:
		return 2
	default:
		return 0
	}
}

// ruthAaron は、整数 n が連続する 2 つの整数のうちどちらかとして
// Ruth–Aaron ペアに該当するかどうかを判定して返します。
//   - もしペアに該当しないなら bool の false
//   - ペアに該当するなら []interface{}{ "Ruth" or "Aaron", 1 or 2 or 3 }
func ruth_aaron(n int) interface{} {
	// (n, n+1) を調べる (n が小さい方の場合)
	pairType := checkPair(n, n+1)
	if pairType != 0 {
		return []interface{}{"Ruth", pairType}
	}
	// (n-1, n) を調べる (n が大きい方の場合)
	pairType = checkPair(n-1, n)
	if pairType != 0 {
		return []interface{}{"Aaron", pairType}
	}

	return false

}

// func main() {
// 	fmt.Println(ruth_aaron(5))     // -> ['Ruth',3])
// 	fmt.Println(ruth_aaron(25))    // -> ['Aaron',1])
// 	fmt.Println(ruth_aaron(498))   // -> False)
// 	fmt.Println(ruth_aaron(125))   // -> ['Ruth',2])
// 	fmt.Println(ruth_aaron(715))   // -> ['Aaron',3])
// 	fmt.Println(ruth_aaron(1470))  // -> False)
// 	fmt.Println(ruth_aaron(21183)) // -> ['Ruth',1])
// 	fmt.Println(ruth_aaron(5561))  // -> ['Aaron',2])
// 	fmt.Println(ruth_aaron(6225))  // -> False)
// 	fmt.Println(ruth_aaron(26642)) // -> ['Ruth',3])
// 	fmt.Println(ruth_aaron(18656)) // -> ['Aaron',1])
// 	fmt.Println(ruth_aaron(8558))  // -> False)

// }
