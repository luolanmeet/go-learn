package main

import "algorithm/dataStructure"

func pathSum(root *dataStructure.TreeNode, targetSum int) (ans [][]int) {

	res := [][]int{}
	paths := []int{}

	var dfs func(*dataStructure.TreeNode, int)

	dfs = func(node *dataStructure.TreeNode, left int) {

		if node == nil {
			return
		}
		left -= node.Val

		paths = append(paths, node.Val)
		defer func() {
			paths = paths[:len(paths)-1]
		}()

		if node.Left == nil && node.Right == nil && left == 0 {
			res = append(res, append([]int(nil), paths...))
		}

		dfs(node.Left, left)
		dfs(node.Right, left)
	}

	return res
}

func pathSum2(root *dataStructure.TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	return findPath(root, [][]int{}, []int{}, 0, targetSum)
}

func findPath(root *dataStructure.TreeNode, res [][]int, paths []int, nowSum int, targetSum int) [][]int {

	// 直接用paths，不用copy，会有坑
	paths = append(paths, root.Val)
	var tmp = make([]int, len(paths))
	copy(tmp, paths)

	if root.Left == nil && root.Right == nil {
		if root.Val+nowSum == targetSum {
			//res = append(res, append(paths, root.Val))
			res = append(res, paths)
		}
		return res
	}

	if root.Left != nil {
		res = findPath(root.Left, res, tmp, nowSum+root.Val, targetSum)
	}

	if root.Right != nil {
		res = findPath(root.Right, res, tmp, nowSum+root.Val, targetSum)
	}

	return res
}

func main() {

	// LC中等题 [113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/)

	node13 := dataStructure.TreeNode{Val: 13}
	node7 := dataStructure.TreeNode{Val: 7}
	node2 := dataStructure.TreeNode{Val: 2}
	node5 := dataStructure.TreeNode{Val: 5}
	node1 := dataStructure.TreeNode{Val: 1}
	node11 := dataStructure.TreeNode{Val: 11, Left: &node7, Right: &node2}
	nodeL4 := dataStructure.TreeNode{Val: 4, Left: &node11}
	nodeR4 := dataStructure.TreeNode{Val: 4, Left: &node5, Right: &node1}
	node8 := dataStructure.TreeNode{Val: 8, Left: &node13, Right: &nodeR4}
	root := dataStructure.TreeNode{Val: 5, Left: &nodeL4, Right: &node8}

	pathSum2(&root, 22) // 这个自己写的，不对
	pathSum(&root, 22)

}
