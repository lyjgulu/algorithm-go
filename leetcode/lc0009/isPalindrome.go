package leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return false
	}
	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i <= l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
