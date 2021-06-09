package leetcode

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < l; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
