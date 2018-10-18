package main

import "fmt"

func lengthOfNoRepeatingSubStr (s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch];
			ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func test(s string) int {
	dirct := make(map[byte]int)
	strart := 0
	max := 0
	for i, ch := range []byte(s) {
		if v,ok := dirct[ch];
			ok && v >= strart {
			strart = dirct[ch]
		}

		if i - strart + 1 > max {
			max = i - strart + 1
		}
		dirct[ch] = i
	}
	return max
}

func main() {
	fmt.Println(lengthOfNoRepeatingSubStr("abc"))
	fmt.Println(test("abc"))
}
