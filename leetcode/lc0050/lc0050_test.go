package leetcode

import (
	"fmt"
	"testing"
)

type question46 struct {
	para46
	ans46
}

// para 是参数
// one 代表第一个参数
type para46 struct {
	one float64
	two int
}

// ans 是答案
// one 代表第一个答案
type ans46 struct {
	ans float64
}

func TestProblem46(t *testing.T) {

	qs := []question46{
		{
			para46{5, 3},
			ans46{125},
		},

		{
			para46{5, -3},
			ans46{1.0 / 125},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 46------------------------\n")

	for _, q := range qs {
		a, p := q.ans46, q.para46
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, myPow(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
