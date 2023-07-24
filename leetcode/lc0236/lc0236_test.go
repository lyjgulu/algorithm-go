package leetcode

import (
	"fmt"
	"testing"

	"github.com/lyjgulu/algorithm-go/structures"
)

type question236 struct {
	para236
	ans236
}

// para 是参数
// one 代表第一个参数
type para236 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans236 struct {
	one []int
}

// TODO
func Test_Problem236(t *testing.T) {

	qs := []question236{

		{
			para236{[]int{4, 2, 7, 1, 3, 6, 9}},
			ans236{[]int{4, 7, 2, 9, 6, 3, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 236------------------------\n")

	for _, q := range qs {
		_, p := q.ans236, q.para236
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", structures.Tree2ints(invertTree(root)))
	}
	fmt.Printf("\n\n\n")
}
