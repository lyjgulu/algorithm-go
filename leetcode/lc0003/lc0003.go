package leetcode

// 位图法
func lengthOfLongestSubstring1(s string) int {
	if len(s) <= 0 {
		return 0
	}
	bitSet := [256]bool{}
	left, right, result := 0, 0, 0
	for left < len(s) {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 滑动窗口
func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 0 {
		return 0
	}
	req := [128]int{}
	left, right, result := 0, -1, 0
	for left < len(s) {
		if right+1 < len(s) && req[s[right+1]] == 0 {
			req[s[right+1]]++
			right++
		} else {
			req[s[left]]--
			left++
		}
		if result < right-left+1 {
			result = right - left + 1
		}
	}
	return result
}

// 解法三 滑动窗口-哈希桶
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	window := make(map[byte]int, len(s))
	left, right, result := 0, 0, 0
	for left < len(s) {
		if index, ok := window[s[left]]; ok && index >= right {
			right = index + 1
		}
		window[s[left]] = left
		left++
		if result < left-right {
			result = left - right
		}
	}
	return result
}
