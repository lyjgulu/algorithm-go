package algorithm

func maxArea(height []int) int {
	l := len(height)
	n1, n2 := 0, l-1
	max := 0
	for n1 < n2 {
		area := 0
		if height[n1] >= height[n2] {
			area = (n2 - n1) * height[n2]
			n2--
		}
		if height[n1] < height[n2] {
			area = (n2 - n1) * height[n1]
			n1++
		}
		if area > max {
			max = area
		}
	}
	return max
}
