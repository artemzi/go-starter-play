// Package charrepl solve:
// https://leetcode.com/problems/longest-repeating-character-replacement/description/
/*
Given a string that consists of only uppercase English letters,
you can replace any letter in the string with another letter at most k times.
Find the length of a longest substring containing all repeating letters
you can get after performing the above operations.

Note:
Both the string's length and k will not exceed 10^4.
*/
package charrepl

func characterReplacement(s string, k int) int {
	var (
		max, l int
	)
	count := make([]int, 128)

	for r := 0; r < len(s); r++ {
		count[s[r]]++
		if max < count[s[r]] {
			max = count[s[r]]
		}

		if r-l+1-max > k {
			count[s[l]]--
			l++
		}
	}

	return len(s) - l
}

/*
result: You are here!
Your runtime beats 100.00 % of golang submissions.
(details.png)

solution:
	The whole idea is that if we have a string of length N out of which M characters are same,
	we can replace (N - M) characters to get a continueous string of N characters.
	If M <= K. N is the local maximum for this window.
	If this length is greater than K. Slide the window.
*/
