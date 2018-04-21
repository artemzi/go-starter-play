// Most Common Word
// https://leetcode.com/problems/most-common-word/description/

// Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.
// It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

// Words in the list of banned words are given in lowercase, and free of punctuation.
// Words in the paragraph are not case sensitive.  The answer is in lowercase.

/******
46 / 46 test cases passed.
Status: Accepted
Runtime: 8 ms
Submitted: 0 minutes ago
*/

// Package common must have comment
package common

import (
	"fmt"
	"strings"
)

const delim = "?!.;,*'"

func isDelim(c string) bool {
	if strings.Contains(delim, c) {
		return true
	}
	return false
}

func cleanString(input string) string {

	size := len(input)
	temp := ""
	var prevChar string

	for i := 0; i < size; i++ {
		str := string(input[i]) // convert to string for easier operation
		if (str == " " && prevChar != " ") || !isDelim(str) {
			temp += str
			prevChar = str
		} else if prevChar != " " && isDelim(str) {
			temp += ""
		}
	}
	return strings.ToLower(temp)
}

func contains(m []string, s string) bool {
	for _, a := range m {
		if a == s {
			return true
		}
	}
	return false
}

func mostCommonWord(paragraph string, banned []string) string {
	var result string
	data := strings.Split(cleanString(paragraph), " ")
	c := make(map[string]int)

	for _, w := range data {
		// skip if word in banned list
		if err := contains(banned, w); err {
			continue
		}

		_, ok := c[w]
		if !ok {
			c[w] = 1
			continue
		}
		c[w]++
	}

	fmt.Printf("%v\n", c)

	tmp := 0
	for k, v := range c {
		if tmp == 0 {
			tmp = v
			result = k
			continue
		}

		if tmp < v {
			tmp = v
			result = k
			continue
		}
	}

	return result
}

func test() {
	line := "j. t? T. z! R, v, F' x! L; l! W. M; S. y? r! n; O. q; I? h; w. t; y; X? y, p. k! k, h, J, r? w! U! V; j' u; R! z. s. T' k. P? M' I' j! y. P, T! e; X. w? M! Y, X; G; d, X? S' F, K? V, r' v, v, D, w, K! S? Q! N. n. V. v. t? t' x! u. j; m; n! F, V' Y! h; c! V, v, X' X' t? n; N' r; x. W' P? W; p' q, S' X, J; R. x; z; z! G, U; m. P; o. P! Y; I, I' l' J? h; Q; s? U, q, x. J, T! o. z, N, L; u, w! u, S. Y! V; S? y' E! O; p' X, w. p' M, h! R; t? K? Y' z? T? w; u. q' R, q, T. R? I. R! t, X, s? u; z. u, Y, n' U; m; p? g' P? y' v, o? K? R. Q? I! c, X, x. r' u! m' y. t. W; x! K? B. v; m, k; k' x; Z! U! p. U? Q, t, u' E' n? S' w. y; W, x? r. p! Y? q, Y. t, Z' V, S. q; W. Z, z? x! k, I. n; x? z; V? s! g, U; E' m! Z? y' x? V! t, F. Z? Y' S! z, Y' T? x? v? o! l; d; G' L. L, Z? q. w' r? U! E, H. C, Q! O? w! s? w' D. R, Y? u. w, N. Z? h. M? o, B, g, Z! t! l, W? z, o? z, q! O? u, N; o' o? V; S! z; q! q. o, t! q! w! Z? Z? w, F? O' N' U' p? r' J' L; S. M; g' V. i, P, v, v, f; W? L, y! i' z; L? w. v, s! P?"
	ban := []string{"m", "q", "e", "l", "c", "i", "z", "j", "g", "t", "w", "v", "h", "p", "d", "b", "a", "r", "x", "n"}

	fmt.Println(mostCommonWord(line, ban))
}
