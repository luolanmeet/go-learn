package main

import (
	"fmt"
)

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	dp1, dp2 := 0, nums[0]
	for i := 2; i <= len(nums); i++ {

		tmp := dp1 + nums[i-1]
		if dp2 > tmp {
			tmp = dp2
		}
		dp1 = dp2
		dp2 = tmp
	}

	return dp2
}

func main() {
	// [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
