package main

import (
	"algorithm/dataStructure"
	"fmt"
)

func lowestCommonAncestor(root, p, q *dataStructure.TreeNode) *dataStructure.TreeNode {

	var pPaths []*dataStructure.TreeNode
	var qPaths []*dataStructure.TreeNode
	paths := []*dataStructure.TreeNode{}

	var dfs func(node *dataStructure.TreeNode)
	dfs = func(node *dataStructure.TreeNode) {

		if node == nil {
			return
		}
		if len(pPaths) > 0 && len(qPaths) > 0 {
			return
		}

		paths = append(paths, node)
		defer func() {
			paths = paths[:len(paths)-1]
		}()

		if node.Val == p.Val {
			pPaths = make([]*dataStructure.TreeNode, len(paths))
			copy(pPaths, paths)
		}
		if node.Val == q.Val {
			qPaths = make([]*dataStructure.TreeNode, len(paths))
			copy(qPaths, paths)
		}

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	if len(qPaths) > len(pPaths) {
		qPaths, pPaths = pPaths, qPaths
	}

	var res *dataStructure.TreeNode

	for i := range qPaths {
		if qPaths[i].Val == pPaths[i].Val {
			res = qPaths[i]
		}
	}

	return res
}

func main() {
	// [236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)

	root := initTree()
	res := lowestCommonAncestor(root, root.Right, root.Right.Right.Left)
	fmt.Println(res.Val)
}

func initTree() *dataStructure.TreeNode {

	/**
	 *     1
	 *   2   3
	 * 4       5
	 *        6
	 *      7   8
	 */

	t8 := dataStructure.TreeNode{Val: 8}
	t7 := dataStructure.TreeNode{Val: 7}
	t6 := dataStructure.TreeNode{Val: 6, Left: &t7, Right: &t8}
	t5 := dataStructure.TreeNode{Val: 5, Left: &t6}
	t4 := dataStructure.TreeNode{Val: 4}
	t3 := dataStructure.TreeNode{Val: 3, Right: &t5}
	t2 := dataStructure.TreeNode{Val: 2, Left: &t4}
	t1 := dataStructure.TreeNode{Val: 1, Left: &t2, Right: &t3}

	return &t1
}
