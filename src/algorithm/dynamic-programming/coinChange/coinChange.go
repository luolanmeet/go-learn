package main

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func main() {
	// [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)
}
