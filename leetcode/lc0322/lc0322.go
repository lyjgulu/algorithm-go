package leetcode

import "math"

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		if memo[n] != 0 {
			return memo[n]
		}
		res := math.MaxInt32
		for _, coin := range coins {
			subProblem := dp(n - coin)
			if subProblem == -1 {
				continue
			}
			res = int(math.Min(float64(res), float64(1+subProblem)))
		}
		if res != math.MaxInt32 {
			memo[n] = res
		} else {
			memo[n] = -1
		}
		return memo[n]
	}
	return dp(amount)
}

func coinChange1(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
