package leetcode

import "sort"

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	var res [][]int
	var dfs func([]int, int)
	dfs = func(track []int, index int) {
		if index == l {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
			return
		}
		track = append(track, nums[index])
		dfs(track, index+1)
		track = track[:len(track)-1]
		dfs(track, index+1)
	}
	dfs([]int{}, 0)
	return res
}
