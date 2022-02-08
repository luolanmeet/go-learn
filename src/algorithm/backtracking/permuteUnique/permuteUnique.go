package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {

	res := [][]int{}
	combo := []int{}
	used := make([]bool, len(nums))

	sort.Ints(nums)

	var dfs func()

	dfs = func() {

		if len(combo) == len(nums) {
			res = append(res, append([]int(nil), combo...))
			return
		}

		preVal := -100
		for i := 0; i < len(nums); i++ {

			if used[i] || preVal == nums[i] {
				continue
			}

			used[i] = true
			combo = append(combo, nums[i])

			dfs()
			used[i] = false
			combo = combo[:len(combo)-1]
			preVal = nums[i]
		}
	}

	dfs()

	return res
}

func main() {
	// [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)
	res := permuteUnique([]int{1, 1, 2})
	fmt.Printf("%v", res)
}
