// Two Sum code problem
// https://leetcode.com/problems/reverse-integer/description/
package main

import (
	"fmt"
	"strconv"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse(x int) int {
	var (
		res      int
		s        = strconv.Itoa(x)          // get string from passed int
		reversed = []rune(reverseString(s)) // reverse string and convert to rune
	)

	// check minus sign
	if reversed[len(reversed)-1] == '-' {
		res, _ = strconv.Atoi("-" + string(reversed[:len(reversed)-1]))
	} else {
		res, _ = strconv.Atoi(string(reversed))
	}

	// check int32 overflow
	if res < -2147483648 || res > 2147483647 {
		return 0
	}

	return res
}

func main() {
	fmt.Printf("123 => %d\n", reverse(123))
	fmt.Printf("-123 => %d\n", reverse(-123))
}
