package leetcode

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, val := range s {
		if (val == '[') || (val == '{') || (val == '(') {
			stack = append(stack, val)
		} else if ((val == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((val == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((val == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
