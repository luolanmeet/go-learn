package main

import (
	"algorithm/dataStructure"
	"fmt"
)

/**
 * 前序遍历——迭代
 */
func pre(head *dataStructure.TreeNode) {

	if head == nil {
		return
	}
	stack := []*dataStructure.TreeNode{head}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(node.Val, " ")
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

/**
 * 前序遍历——递归
 */
func preRecursion(node *dataStructure.TreeNode) {

	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	preRecursion(node.Left)
	preRecursion(node.Right)
}

/**
 * 中序遍历——迭代
 */
func mid(node *dataStructure.TreeNode) {

	if node == nil {
		return
	}

	var stack []*dataStructure.TreeNode
	for node != nil || len(stack) > 0 {

		if node != nil {
			stack = append(stack, node)
			node = node.Left
			continue
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(node.Val, " ")
		node = node.Right
	}

}

/**
 * 中序遍历——递归
 */
func midRecursion(node *dataStructure.TreeNode) {

	if node == nil {
		return
	}
	midRecursion(node.Left)
	fmt.Print(node.Val, " ")
	midRecursion(node.Right)
}

/**
 * 后续遍历——迭代
 */
func post(head *dataStructure.TreeNode) {

	if head == nil {
		return
	}

	// 改造前序遍历（根左右），求出根右左；在反过来遍历，变为左右根

	stack1 := []*dataStructure.TreeNode{head}
	var stack2 []*dataStructure.TreeNode

	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node)
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	for i := len(stack2) - 1; i >= 0; i-- {
		fmt.Print(stack2[i].Val, " ")
	}

}

/**
 * 后续遍历——递归
 */
func postRecursion(node *dataStructure.TreeNode) {
	if node == nil {
		return
	}
	postRecursion(node.Left)
	postRecursion(node.Right)
	fmt.Print(node.Val, " ")
}

/**
 * 层次遍历
 */
func level(node *dataStructure.TreeNode) {

	if node == nil {
		return
	}

	queue := []*dataStructure.TreeNode{node}

	for len(queue) > 0 {

		size := len(queue)
		for i := 0; i < size; i++ {

			node = queue[i]

			fmt.Print(node.Val, " ")
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
	}

}

func main() {
	root := initTree()
	fmt.Println("前序遍历-根左右")
	pre(root)
	fmt.Println()
	preRecursion(root)
	fmt.Println()

	fmt.Println("中序遍历-左根右")
	mid(root)
	fmt.Println()
	midRecursion(root)
	fmt.Println()

	fmt.Println("后续遍历-左右根")
	post(root)
	fmt.Println()
	postRecursion(root)
	fmt.Println()

	fmt.Println("层次遍历遍历")
	level(root)

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
