package leetcode

import (
	"fmt"
	"github.com/lyjgulu/algorithm/structures"
	"testing"
)

type question98 struct {
	para98
	ans98
}

// para 是参数
// one 代表第一个参数
type para98 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans98 struct {
	one []int
}

func Test_Problem98(t *testing.T) {

	qs := []question98{

		{
			para98{[]int{}},
			ans98{[]int{}},
		},

		{
			para98{[]int{1}},
			ans98{[]int{1}},
		},

		{
			para98{[]int{1, structures.NULL, 2, 3}},
			ans98{[]int{1, 3, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 98------------------------\n")

	for _, q := range qs {
		_, p := q.ans98, q.para98
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", inorderTraversal(root))
	}
	fmt.Printf("\n\n\n")
}
