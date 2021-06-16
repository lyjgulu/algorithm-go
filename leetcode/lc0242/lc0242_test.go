package leetcode

import (
	"fmt"
	"testing"
)

type question242 struct {
	para242
	ans242
}

// para 是参数
// one 代表第一个参数
type para242 struct {
	s string
	t string
}

// ans 是答案
// one 代表第一个答案
type ans242 struct {
	one bool
}

func TestProblem242(t *testing.T) {

	qs := []question242{
		{
			para242{"anagram", "nagaram"},
			ans242{true},
		},

		{
			para242{"rat", "car"},
			ans242{false},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 242------------------------\n")

	for _, q := range qs {
		a, p := q.ans242, q.para242
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, isAnagram(p.s, p.t))
	}
	fmt.Printf("\n\n\n")
}
