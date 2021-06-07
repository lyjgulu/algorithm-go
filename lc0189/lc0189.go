package lc0189

func rotate(nums []int, k int) []int {
	l := len(nums)
	if l == 0 || l == 1 {
		return nums
	}
	k %= l
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	return nums
}

func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}
