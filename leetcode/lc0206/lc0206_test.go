package leetcode

import (
	"fmt"
	"testing"

	"github.com/lyjgulu/algorithm-go/structures"
)

type question206 struct {
	para206
	ans206
}

// para 是参数
// one 代表第一个参数
type para206 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans206 struct {
	one []int
}

func Test_Problem206(t *testing.T) {

	qs := []question206{

		{
			para206{[]int{1, 2, 3, 4, 5}},
			ans206{[]int{5, 4, 3, 2, 1}},
		},

		{
			para206{[]int{1, 2}},
			ans206{[]int{2, 1}},
		},

		{
			para206{[]int{}},
			ans206{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 206------------------------\n")

	for _, q := range qs {
		a, p := q.ans206, q.para206
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, structures.List2Ints(reverseList(structures.Ints2List(p.one))))

	}
}
