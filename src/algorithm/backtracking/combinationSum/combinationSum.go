package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)

	res := [][]int{}
	combo := []int{}
	var back func(int, int)

	back = func(idx int, target int) {

		if target == 0 {
			tmp := make([]int, len(combo))
			copy(tmp, combo)
			res = append(res, tmp)
			return
		}

		if idx >= len(candidates) || target < 0 {
			return
		}

		// 不要
		back(idx+1, target)
		combo = append(combo, candidates[idx])
		back(idx, target-candidates[idx])
		combo = combo[:len(combo)-1]
	}
	back(0, target)

	return res
}

func main() {
	// [39. 组合总和](https://leetcode-cn.com/problems/combination-sum/)
}
