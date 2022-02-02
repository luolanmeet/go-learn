package main

import "algorithm/dataStructure"

func rightSideView(root *dataStructure.TreeNode) []int {

	res := []int{}

	if root == nil {
		return res
	}

	queue := []*dataStructure.TreeNode{root}

	size := len(queue)
	for size > 0 {

		res = append(res, queue[size-1].Val)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		size = len(queue)
	}

	return res
}

func main() {
	// [199. 二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/)
}
