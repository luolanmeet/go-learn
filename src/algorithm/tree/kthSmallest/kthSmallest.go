package main

import "algorithm/dataStructure"

func kthSmallest(root *dataStructure.TreeNode, k int) int {

	var stack []*dataStructure.TreeNode

	for root != nil || len(stack) > 0 {

		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}

		k--
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}

	return -1
}

func main() {
	// [230. 二叉搜索树中第K小的元素](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/submissions/)
}
