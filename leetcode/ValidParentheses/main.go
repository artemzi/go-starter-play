// Two Sum code problem
// https://leetcode.com/problems/valid-parentheses/description/
package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	for {
		if strings.Contains(s, "()") {
			s = strings.Replace(s, "()", "", -1)
			continue
		}
		if strings.Contains(s, "{}") {
			s = strings.Replace(s, "{}", "", -1)
			continue
		}
		if strings.Contains(s, "[]") {
			s = strings.Replace(s, "[]", "", -1)
			continue
		} else {
			break
		}
	}

	return (s == "")
}

func main() {
	fmt.Printf("() valid? > %v\n", isValid("()"))
	fmt.Printf("()[]{} valid? > %v\n", isValid("()[]{}"))
	fmt.Printf("(] valid? > %v\n", isValid("(]"))
	fmt.Printf("([)] valid? > %v\n", isValid("([)]"))
	fmt.Printf("(([]){}) valid? > %v\n", isValid("(([]){})"))
}
