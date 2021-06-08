package leetcode

func moveZeroes(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	n1, n2 := 0, 0
	for n2 < l {
		if nums[n2] != 0 {
			nums[n1], nums[n2] = nums[n2], nums[n1]
			n1++
		}
		n2++
	}
	return nums
}
