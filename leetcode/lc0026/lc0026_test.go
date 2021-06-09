package leetcode

import (
	"fmt"
	"testing"
)

type question26 struct {
	para26
	ans26
}

// para 是参数
// one 代表第一个参数
type para26 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans26 struct {
	one int
}

func TestProblem26(t *testing.T) {

	qs := []question26{
		{
			para26{[]int{1, 1, 2}},
			ans26{2},
		},

		{
			para26{[]int{-1, -100, 3, 99}},
			ans26{4},
		},

		{
			para26{[]int{3, 2, 4}},
			ans26{3},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 26------------------------\n")

	for _, q := range qs {
		a, p := q.ans26, q.para26
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, removeDuplicates(p.nums))
	}
	fmt.Printf("\n\n\n")
}
