package leetcode

import (
	"fmt"
	"testing"
)

type question11 struct {
	para11
	ans11
}

// para 是参数
// one 代表第一个参数
type para11 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans11 struct {
	ans int
}

func TestProblem11(t *testing.T) {

	qs := []question11{
		{
			para11{[]int{1, 1}},
			ans11{1},
		},

		{
			para11{[]int{4, 3, 2, 1, 4}},
			ans11{16},
		},

		{
			para11{[]int{1, 2, 1}},
			ans11{2},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 11------------------------\n")

	for _, q := range qs {
		a, p := q.ans11, q.para11
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, maxArea(p.nums))
	}
	fmt.Printf("\n\n\n")
}
