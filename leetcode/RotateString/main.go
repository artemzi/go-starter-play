// Package rotatestring solve leetcode problem
// https://leetcode.com/problems/rotate-string/description/
package rotatestring

func move(s *string) {
	l := len((*s)) - 1
	r := []rune((*s))
	r = append([]rune{r[l]}, r[:l]...)
	(*s) = string(r)
}

func rotateString(A string, B string) bool {

	return false
}
