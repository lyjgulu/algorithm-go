package leetcode

import (
	"fmt"
	"github.com/lyjgulu/algorithm/structures"
	"testing"
)

type question124 struct {
	para124
	ans124
}

// para 是参数
// one 代表第一个参数
type para124 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans124 struct {
	one int
}

func Test_Problem124(t *testing.T) {

	qs := []question124{

		{
			para124{[]int{-10, 9, 20, 0, 0, 15, 7}},
			ans124{42},
		},

		{
			para124{[]int{1, 2, 3, 4, 5}},
			ans124{11},
		},

		{
			para124{[]int{1, 2}},
			ans124{3},
		},

		{
			para124{[]int{}},
			ans124{-2147483648},
		},

		{
			para124{[]int{1}},
			ans124{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 124------------------------\n")

	for _, q := range qs {
		a, p := q.ans124, q.para124
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, maxPathSum(structures.Ints2TreeNode(p.one)))

	}
}
