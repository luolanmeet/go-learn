package main

import "algorithm/dataStructure"

func sumNumbers(root *dataStructure.TreeNode) int {

	res := 0
	if root == nil {
		return 0
	}

	var dfs func(root *dataStructure.TreeNode, nowNum int)
	dfs = func(node *dataStructure.TreeNode, nowNum int) {

		nowNum = nowNum*10 + node.Val

		if node.Left == nil && node.Right == nil {
			res = res + nowNum
			return
		}

		if node.Left != nil {
			dfs(node.Left, nowNum)
		}
		if node.Right != nil {
			dfs(node.Right, nowNum)
		}
	}

	dfs(root, 0)

	return res
}

func main() {
	// [129. 求根到叶子节点数字之和](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/)
}
