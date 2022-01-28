package main

import "algorithm/dataStructure"

func buildTree(inorder []int, postorder []int) *dataStructure.TreeNode {

	size := len(postorder)
	if size == 0 {
		return nil
	}

	node := dataStructure.TreeNode{Val: postorder[size-1]}
	if size == 1 {
		return &node
	}

	idx := 0
	for i := range inorder {
		if inorder[i] == postorder[size-1] {
			idx = i
			break
		}
	}

	// 递归时，可以通过检查  inorder 切片和 postorder切片的长度是否相等，来简单检查
	node.Left = buildTree(inorder[0:idx], postorder[0:idx])
	if idx < size-1 {
		node.Right = buildTree(inorder[idx+1:], postorder[idx:size-1])
	}
	return &node
}

func main() {
	// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
	// `LC中等题`[106. 从中序与后序遍历序列构造二叉树]
}
