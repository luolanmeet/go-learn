package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	var res [][]int
	var combo []int

	sort.Ints(candidates)

	var dfs func(int, int)
	dfs = func(target int, idx int) {

		if target == 0 {
			res = append(res, append([]int(nil), combo...))
			return
		}

		if idx >= len(candidates) || target < 0 {
			return
		}

		combo = append(combo, candidates[idx])
		dfs(target-candidates[idx], idx+1)
		combo = combo[:len(combo)-1]

		for i := idx + 1; i < len(candidates); i++ {
			if candidates[i] == candidates[idx] {
				continue
			}
			dfs(target, i)
			break
		}
	}

	dfs(target, 0)

	return res
}

func main() {
	// [40. 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/)
}
