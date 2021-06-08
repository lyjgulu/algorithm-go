package leetcode

import (
	"fmt"
	"testing"
)

type question283 struct {
	para283
	ans283
}

// para 是参数
// one 代表第一个参数
type para283 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans283 struct {
	one []int
}

func TestProblem283(t *testing.T) {

	qs := []question283{
		{
			para283{[]int{0, 1, 0, 3, 12}},
			ans283{[]int{1, 3, 12, 0, 0}},
		},

		{
			para283{[]int{1, 1, 0, 3, 12}},
			ans283{[]int{1, 1, 3, 12, 0}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 283------------------------\n")

	for _, q := range qs {
		a, p := q.ans283, q.para283
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, moveZeroes(p.nums))
	}
	fmt.Printf("\n\n\n")
}
