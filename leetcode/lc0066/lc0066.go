package leetcode

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits[:]...)
}
