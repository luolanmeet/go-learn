package main

import (
	"algorithm/dataStructure"
	"fmt"
)

func isSymmetric(root *dataStructure.TreeNode) bool {

	if root == nil {
		return true
	}

	// 层次遍历
	queue := []*dataStructure.TreeNode{root}

	for len(queue) > 0 {

		size := len(queue)
		if size != 1 && size%2 != 0 {
			return false
		}

		var tmpInt []int

		for i := 0; i < size; i++ {

			node := queue[i]
			// 暂存同一层节点的数值
			tmpInt = append(tmpInt, node.Val)

			if node.Val == 1000 {
				continue
			}

			// 用 1000 表示nil节点

			if node.Left != nil {
				queue = append(queue, node.Left)
			} else {
				queue = append(queue, &dataStructure.TreeNode{Val: 1000})
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			} else {
				queue = append(queue, &dataStructure.TreeNode{Val: 1000})
			}
		}

		queue = queue[size:]

		for i := 0; i < size/2; i++ {
			if tmpInt[i] != tmpInt[size-1-i] {
				return false
			}
		}
	}

	return true
}

func main() {

	// 101. 是否是对称二叉树
	// 轴对称
	// https://leetcode-cn.com/problems/symmetric-tree/

	node := initTree()
	fmt.Println(isSymmetric(node))
}

func initTree() *dataStructure.TreeNode {

	/**
	 *       1
	 *    2     2
	 *  3  4  4   3
	 */

	t7 := dataStructure.TreeNode{Val: 3}
	t6 := dataStructure.TreeNode{Val: 4}
	t5 := dataStructure.TreeNode{Val: 2, Left: &t7, Right: &t6}

	t4 := dataStructure.TreeNode{Val: 4}
	t3 := dataStructure.TreeNode{Val: 3}
	t2 := dataStructure.TreeNode{Val: 2, Left: &t4, Right: &t3}

	t1 := dataStructure.TreeNode{Val: 1, Left: &t5, Right: &t2}

	return &t1
}
