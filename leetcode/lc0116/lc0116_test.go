package leetcode

import (
	"fmt"
	"testing"

	"github.com/lyjgulu/algorithm-go/structures"
)

type question116 struct {
	para116
	ans116
}

// para 是参数
// one 代表第一个参数
type para116 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans116 struct {
	one int
}

func Test_Problem116(t *testing.T) {

	qs := []question116{

		{
			para116{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans116{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 116------------------------\n")

	for _, q := range qs {
		_, p := q.ans116, q.para116
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", maxDepth(root))
	}
	fmt.Printf("\n\n\n")
}
