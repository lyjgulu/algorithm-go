package leetcode

import (
	"fmt"
	"testing"

	"github.com/lyjgulu/algorithm-go/structures"
)

type question21 struct {
	para21
	ans21
}

// para 是参数
// one 代表第一个参数
type para21 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans21 struct {
	one []int
}

func Test_Problem21(t *testing.T) {

	qs := []question21{

		{
			para21{[]int{1, 2, 4}, []int{1, 3, 4}},
			ans21{[]int{1, 1, 2, 3, 4, 4}},
		},

		{
			para21{[]int{}, []int{}},
			ans21{[]int{}},
		},

		{
			para21{[]int{}, []int{0}},
			ans21{[]int{0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 21------------------------\n")

	for _, q := range qs {
		a, p := q.ans21, q.para21
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, structures.List2Ints(mergeTwoLists(structures.Ints2List(p.one), structures.Ints2List(p.two))))

	}
}
