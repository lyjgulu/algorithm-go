package algorithm

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func TestProblem1(t *testing.T) {

	qs := []question1{
		{
			para1{[]int{2, 7, 11, 15}, 9},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{-1, -100, 3, 99}, 2},
			ans1{[]int{0, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		a, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, twoSum(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
