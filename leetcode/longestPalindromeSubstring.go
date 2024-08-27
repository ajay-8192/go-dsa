package leetcode

import "fmt"

func LongestPalindromeSubstring(s string) string {
	// babad: bab | aba

	// format string to @b#a#b#a#d#$
	Q := "@"
	stringLength := len(s)

	for i := 0; i < stringLength; i++ {
		Q += "#" + string(s[i])
	}

	Q += "#$"

	P := make([]int, len(Q)+1)

	c := 0
	r := 0
	for i := 0; i < len(Q)-1; i++ {

		iMirror := c - (i - c)

		if r > i {
			P[i] = min(r-i, P[iMirror])
		}

		a := i + 1 + P[i]
		b := i - 1 - P[i]
		for a < len(Q) && b >= 0 && Q[a] == Q[b] {
			P[i] = P[i] + 1
			a++
			b--
		}

		if i+P[i] > r {
			c = i
			r = i + P[i]
		}
	}

	maxPalindrome := 0
	centerIndex := 0

	for i := 0; i < len(Q)-1; i++ {
		if P[i] > maxPalindrome {
			maxPalindrome = P[i]
			centerIndex = i
		}
	}
	fmt.Println(centerIndex/2, maxPalindrome, s)
	if maxPalindrome%2 == 0 {
		return s[(centerIndex-1)/2-(maxPalindrome-1) : (centerIndex-1)/2+(maxPalindrome-1)]
	}
	return s[(centerIndex-1-maxPalindrome)/2 : maxPalindrome]
}