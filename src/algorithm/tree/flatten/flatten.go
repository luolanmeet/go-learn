package main

import "algorithm/dataStructure"

func flatten(root *dataStructure.TreeNode) {

	if root == nil {
		return
	}

	stack := []*dataStructure.TreeNode{root}
	preNode := &dataStructure.TreeNode{Val: -1}

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		preNode.Left = nil
		preNode.Right = node
		preNode = node

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

}

func main() {
	// [114. 二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)
}
