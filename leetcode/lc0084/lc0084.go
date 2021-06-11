package leetcode

func largestRectangleArea(heights []int) int {
	l := len(heights)
	left, right := make([]int, l), make([]int, l)
	stack := make([]int, 0)
	for i := 0; i < l; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := l - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = l
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	max := 0
	for i := 0; i < l; i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		if max < area {
			max = area
		}
	}
	return max
}
