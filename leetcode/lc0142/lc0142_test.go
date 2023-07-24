package leetcode

import (
	"fmt"
	"testing"

	"github.com/lyjgulu/algorithm-go/structures"
)

type question142 struct {
	para142
	ans142
}

// para 是参数
// one 代表第一个参数
type para142 struct {
	one  []int
	para int
}

// ans 是答案
// one 代表第一个答案
type ans142 struct {
	one int
	two int
}

func Test_Problem142(t *testing.T) {

	qs := []question142{

		{
			para142{[]int{3, 2, 0, -4}, 1},
			ans142{1, 2},
		},

		{
			para142{[]int{1, 2}, 0},
			ans142{1, 1},
		},

		{
			para142{[]int{1}, -1},
			ans142{0, 0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 142------------------------\n")

	for _, q := range qs {
		a, p := q.ans142, q.para142
		answer, val := detectCycle(structures.Ints2ListWithCycle(p.one, p.para))
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v,%v\n", p, a, answer, val)
	}
}
