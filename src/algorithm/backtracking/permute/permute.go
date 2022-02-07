package main

import "fmt"

func permute(nums []int) [][]int {

	res := [][]int{}
	combo := []int{}
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {

		if len(combo) == len(nums) {
			res = append(res, append([]int(nil), combo...))
			return
		}

		for i := 0; i < len(nums); i++ {

			if used[i] {
				continue
			}

			used[i] = true
			combo = append(combo, nums[i])
			dfs()

			combo = combo[:len(combo)-1]
			used[i] = false
		}
	}

	dfs()
	return res
}

func main() {
	// [46. 全排列](https://leetcode-cn.com/problems/permutations/)
	combos := permute([]int{1, 2, 3})
	fmt.Printf("%v", combos)
}
