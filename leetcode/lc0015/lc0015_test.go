package leetcode

import (
	"fmt"
	"testing"
)

type question15 struct {
	para15
	ans15
}

// para 是参数
// one 代表第一个参数
type para15 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans15 struct {
	ans [][]int
}

func TestProblem15(t *testing.T) {

	qs := []question15{
		{
			para15{[]int{-1, 0, 1, 2, -1, -4}},
			ans15{[][]int{{-1, -1, 2}, {-1, 0, 1}}},
		},

		{
			para15{[]int{1, 2, -2, -1}},
			ans15{[][]int{}},
		},

		{
			para15{[]int{}},
			ans15{[][]int{}},
		},

		{
			para15{[]int{0}},
			ans15{[][]int{}},
		},

		{
			para15{[]int{0, 5, 2, 4, 6}},
			ans15{[][]int{}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 15------------------------\n")

	for _, q := range qs {
		a, p := q.ans15, q.para15
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, threeSum(p.nums))
	}
	fmt.Printf("\n\n\n")
}
