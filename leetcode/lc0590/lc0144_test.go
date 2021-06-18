package leetcode

import (
	"fmt"
	"github.com/lyjgulu/algorithm/structures"
	"testing"
)

type question144 struct {
	para144
	ans144
}

// para 是参数
// one 代表第一个参数
type para144 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans144 struct {
	one []int
}

func Test_Problem144(t *testing.T) {

	qs := []question144{

		{
			para144{[]int{1, structures.NULL, 2, 3, 4, 5, structures.NULL, 6, 7, structures.NULL, 8, structures.NULL, 9, 10,
				structures.NULL, structures.NULL, 11, structures.NULL, 12, structures.NULL, 13, structures.NULL, structures.NULL, 14}},
			ans144{[]int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1}},
		},

		{
			para144{[]int{1, structures.NULL, 3, 2, 4, structures.NULL, 5, 6}},
			ans144{[]int{5, 6, 3, 2, 4, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 144------------------------\n")

	for _, q := range qs {
		_, p := q.ans144, q.para144
		fmt.Printf("【input】:%v      ", p)
		root := structures.Int2NaryNode(p.one)
		fmt.Printf("【output】:%v      \n", postorder(root))
	}
	fmt.Printf("\n\n\n")
}
