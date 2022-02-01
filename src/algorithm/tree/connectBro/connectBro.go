package main

import "algorithm/dataStructure"

func connect(root *dataStructure.Node) *dataStructure.Node {

	if root == nil {
		return root
	}

	// 层次遍历
	queue := []*dataStructure.Node{root}
	size := len(queue)
	for size > 0 {
		var preNode *dataStructure.Node
		for i := 0; i < size; i++ {
			node := queue[i]
			if preNode != nil {
				preNode.Next = node
				preNode = node
			} else {
				preNode = node
			}
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

	return root
}

func main() {
	// 116. 填充同一层的兄弟节点](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/)
}
