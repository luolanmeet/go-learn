package tree

/**
 * 递归中序遍历
 */
func (node *Node) MidRecursion() {

	if node == nil {
		return
	}
	// 这里不用判断left是不是为空，Go真的6哇
	node.Left.MidRecursion()
	node.Print()
	node.Right.MidRecursion()
}
