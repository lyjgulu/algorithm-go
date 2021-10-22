package leetcode

import (
	"sort"
)

/*result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择*/
func permute(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var backtrack func([]int, int)
	backtrack = func(track []int, first int) {
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
			return
		}
		for i := first; i < len(nums); i++ {
			flag := false
			for _, val := range track {
				if val == nums[i] {
					flag = true
				}
			}
			if flag {
				continue
			}
			track = append(track, nums[i])
			backtrack(track, first)
			track = track[:len(track)-1]
		}
	}
	backtrack([]int{}, 0)
	return res
}
