package template

var nums []int

func dp(n int) {
	for _, i := range nums {
		dp(n - i)
	}
}
