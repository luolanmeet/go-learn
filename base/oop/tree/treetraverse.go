package tree

import (
	"fmt"
)

/**
 * 递归中序遍历
 */
func (node *Node) MidRecursion() {

	if node == nil {
		return
	}
	// 这里不用判断left是不是为空，Go真的6哇
	node.left.MidRecursion()
	node.print()
	node.right.MidRecursion()
}
