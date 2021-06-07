package algorithm

import "sort"

// 双指针加排序
func threeSum(nums []int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	for first := 0; first < l; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		if nums[first] > 0 {
			break
		}
		third := l - 1
		for second := first + 1; second < third; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for third > second+1 && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
