package leetcode

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// TODO
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var res []string
	var backtrack func(int, string)
	backtrack = func(index int, s string) {
		if index == len(digits) {
			res = append(res, s)
		} else {
			digits := string(digits[index])
			letters := phoneMap[digits]
			lettersCount := len(letters)
			for i := 0; i < lettersCount; i++ {
				backtrack(index+1, s+string(letters[i]))
			}
		}
	}
	backtrack(digits, 0, "")
	return res
}
