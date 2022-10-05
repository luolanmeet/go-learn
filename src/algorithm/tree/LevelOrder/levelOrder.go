package main

import (
	"algorithm/dataStructure"
	"fmt"
)

func levelOrder(root *dataStructure.TreeNode) [][]int {

	queue := []*dataStructure.TreeNode{root}
	var res [][]int

	for len(queue) > 0 {

		size := len(queue)
		var level []int
		for i := 0; i < size; i++ {
			node := queue[i]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
		res = append(res, level)
	}

	return res
}

func main() {

	root := initTree()
	res := levelOrder(root)
	for i := range res {
		for j := range res[i] {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println()
	}
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
