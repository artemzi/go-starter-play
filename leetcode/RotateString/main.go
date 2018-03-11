// Package rotatestring solve leetcode problem
// https://leetcode.com/problems/rotate-string/description/
package rotatestring

func move(s *string) {
	r := []rune((*s))       // convert string into runes array
	r = append(r[1:], r[0]) // move
	(*s) = string(r)        // save result
}

func rotateString(A string, B string) bool {
	tmp := A
	for i := 0; i < len(A); i++ {
		if tmp == B {
			return true
		}
		move(&tmp)
	}
	return false
}
