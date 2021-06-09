package leetcode

import (
	"fmt"
	"testing"
)

type question20 struct {
	para20
	ans20
}

// para 是参数
// one 代表第一个参数
type para20 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans20 struct {
	one bool
}

func TestProblem20(t *testing.T) {

	qs := []question20{
		{
			para20{"()[]{}"},
			ans20{true},
		},
		{
			para20{"(]"},
			ans20{false},
		},
		{
			para20{"({[]})"},
			ans20{true},
		},
		{
			para20{"(){[({[]})]}"},
			ans20{true},
		},
		{
			para20{"((([[[{{{"},
			ans20{false},
		},
		{
			para20{"(())]]"},
			ans20{false},
		},
		{
			para20{""},
			ans20{true},
		},
		{
			para20{"["},
			ans20{false},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 20------------------------\n")

	for _, q := range qs {
		a, p := q.ans20, q.para20
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, isValid(p.one))
	}
	fmt.Printf("\n\n\n")
}
