// https://leetcode.com/problems/rotate-string/description/
package main

import "fmt"

// Data is abstraction for test input
type Data struct {
	A string
	B string
}

func rotateString(A string, B string) bool {

	return true
}

func main() {
	data := make([]Data, 0, 3)
	data = append(data, Data{"abcde", "cdeab"}) // true
	data = append(data, Data{"abcde", "abced"}) // false

	fmt.Printf("%v\n", data)
}
