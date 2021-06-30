package leetcode

import "sort"

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
	// 回溯
	var res [][]int

	var backTrack func(path []int, start int)
	backTrack = func(path []int, start int) {

		if len(path) == len(nums) {
			// 把path添加到res中
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(nums); i++ {
			// 如果当前路径包含相同字母，不添加
			flag := false
			for _, v := range path {
				if v == nums[i] {
					flag = true
				}
			}
			if flag {
				continue
			}

			path = append(path, nums[i])
			backTrack(path, start)
			path = path[:len(path)-1]
		}

	}
	backTrack([]int{}, 0)
	return res
}
