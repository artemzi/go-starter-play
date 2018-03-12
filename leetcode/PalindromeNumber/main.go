// https://leetcode.com/problems/palindrome-number/description/
package main

import "fmt"

func isPalindrome(x int) bool {
	var (
		dig, rev int
	)
	for i := x; i > 0; i /= 10 {
		dig = i % 10
		rev = rev*10 + dig
	}

	if x == rev {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(-2147483648)) // expected false
}
