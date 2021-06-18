package leetcode

// 深度优先遍历
func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return res
	}
	var dfs func(path string, left int, right int)
	dfs = func(path string, left int, right int) {
		if left == 0 && right == 0 {
			res = append(res, path)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			dfs(path+"(", left-1, right)
		}
		if right > 0 {
			dfs(path+")", left, right-1)
		}
	}
	dfs("", n, n)
	return res
}
