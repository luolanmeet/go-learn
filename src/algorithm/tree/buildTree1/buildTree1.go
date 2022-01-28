package main

import (
	"algorithm/dataStructure"
)

func buildTree(preorder []int, inorder []int) *dataStructure.TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	node := dataStructure.TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return &node
	}

	// idx 可以当做 左子树的 size用
	idx := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			idx = i
			break
		}
	}

	node.Left = buildTree(preorder[1:idx+1], inorder[0:idx])
	node.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return &node
}

func main() {
	// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)
	// LC中等题 105. 从前序与中序遍历序列构造二叉树

}
