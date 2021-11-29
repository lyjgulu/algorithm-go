package leetcode

import (
	"fmt"
	"github.com/lyjgulu/algorithm/structures"
	"testing"
)

type question654 struct {
	para654
	ans654
}

// para 是参数
// one 代表第一个参数
type para654 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans654 struct {
	one int
}

func Test_Problem654(t *testing.T) {

	qs := []question654{

		{
			para654{[]int{3, 2, 1, 6, 0, 5}},
			ans654{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 654------------------------\n")

	for _, q := range qs {
		_, p := q.ans654, q.para654
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", structures.Tree2Preorder(constructMaximumBinaryTree(p.one)))
	}
	fmt.Printf("\n\n\n")
}
