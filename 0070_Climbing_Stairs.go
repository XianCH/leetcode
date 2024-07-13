package leetcode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 对于每一级台阶`i`，你可以从第`i-1`级台阶爬一步上来，也可以从第`i-2`级台阶爬两步上来。所以，到达第`i`级台阶的方法数就是到达第`i-1`级台阶的方法数和到达第`i-2`级台阶的方法数的和。
