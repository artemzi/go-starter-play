// Package charrepl solve:
// https://leetcode.com/problems/longest-repeating-character-replacement/description/
/*
Given a string that consists of only uppercase English letters,
you can replace any letter in the string with another letter at most k times.
Find the length of a longest substring containing all repeating letters
you can get after performing the above operations.

Note:
Both the string's length and k will not exceed 104.
*/
package charrepl

func characterReplacement(s string, k int) int {
	if ln := len(s); ln == 0 {
		return 0
	}
	return 1
}
