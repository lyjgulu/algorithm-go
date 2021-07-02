package leetcode

import (
	"fmt"
	"testing"
)

type question70 struct {
	para70
	ans70
}

// para 是参数
// one 代表第一个参数
type para70 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans70 struct {
	ans int
}

func TestProblem70(t *testing.T) {

	/*输入：[3,2,3]
	  输出：3

	  输入：[2,2,1,1,1,2,2]
	  输出：2*/

	qs := []question70{
		{
			para70{[]int{3, 2, 3}},
			ans70{3},
		},

		{
			para70{[]int{2, 2, 1, 1, 1, 2, 2}},
			ans70{2},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 70------------------------\n")

	for _, q := range qs {
		a, p := q.ans70, q.para70
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, majorityElement(p.one))
	}
	fmt.Printf("\n\n\n")
}
