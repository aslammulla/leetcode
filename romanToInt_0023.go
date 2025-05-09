package main

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example,
2 is written as II in Roman numeral, just two ones added together.
12 is written as XII, which is simply X + II.
The number 27 is written as XXVII, which is XX + V + II.
*/

func romanToInt(roman string) int {
	total := 0
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var previous int
	for i := len(roman) - 1; i >= 0; i-- {
		value := m[roman[i]]
		if previous > value {
			total -= value
		} else {
			total += value
		}
		previous = value
	}
	return total
}
