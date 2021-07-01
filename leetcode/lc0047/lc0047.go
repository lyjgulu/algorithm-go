package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	var res [][]int
	condition := make([]bool, l)
	var backtrack func([]int, int)
	backtrack = func(track []int, index int) {
		if index == l {
			temp := make([]int, l)
			copy(temp, track)
			res = append(res, temp)
			return
		}
		for i, v := range nums {
			if condition[i] || i > 0 && !condition[i-1] && v == nums[i-1] {
				continue
			}
			track = append(track, v)
			condition[i] = true
			backtrack(track, index+1)
			condition[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack([]int{}, 0)
	return res
}
