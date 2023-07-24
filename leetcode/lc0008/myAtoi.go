package lc0008

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	res, sign, index := 0, 1, 0
	for inx, c := range s {
		if c == ' ' {
			index = inx
			continue
		}
		if inx-1 == index && c == '+' {
			sign = 1
		} else if inx-1 == index && c == '-' {
			sign = -1
		} else if unicode.IsDigit(c) {
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && c > math.MaxInt32%10) {
				return math.MaxInt32
			}
			if res < math.MinInt32 || (res == math.MinInt32/10 && c < math.MinInt32%10) {
				return math.MinInt32
			}
			res = res*10 + sign*int(c-'0')
		} else {
			break
		}
	}
	return res
}
