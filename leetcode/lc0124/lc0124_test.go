package leetcode

import (
	"fmt"
	"github.com/lyjgulu/algorithm/structures"
	"testing"
)

type question24 struct {
	para24
	ans24
}

// para 是参数
// one 代表第一个参数
type para24 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans24 struct {
	one []int
}

func Test_Problem24(t *testing.T) {

	qs := []question24{

		{
			para24{[]int{1, 2, 3, 4, 5}},
			ans24{[]int{2, 1, 4, 3, 5}},
		},

		{
			para24{[]int{1, 2}},
			ans24{[]int{2, 1}},
		},

		{
			para24{[]int{}},
			ans24{[]int{}},
		},

		{
			para24{[]int{1}},
			ans24{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 24------------------------\n")

	for _, q := range qs {
		a, p := q.ans24, q.para24
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, structures.List2Ints(swapPairs(structures.Ints2List(p.one))))

	}
}
