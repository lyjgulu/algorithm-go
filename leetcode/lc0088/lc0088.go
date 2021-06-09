package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	n1, n2, tail := m-1, n-1, m+n-1
	for ; tail >= 0; tail-- {
		var val int
		if n1 == -1 {
			val = nums2[n2]
			n2--
		} else if n2 == -1 {
			val = nums1[n1]
			n1--
		} else if nums1[n1] > nums2[n2] {
			val = nums1[n1]
			n1--
		} else {
			val = nums2[n2]
			n2--
		}
		nums1[tail] = val
	}
	return nums1
}
