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

/**
 * 利用函数式编程 做遍历时的处理
 */
func (node *Node) TraverseWithFunc(f func(*Node)) {

	if node == nil {
		return
	}

	node.Left.TraverseWithFunc(f)
	f(node)
	node.Right.TraverseWithFunc(f)
}

/**
 * 使用channel返回遍历到的节点
 */
func (node *Node) TraverseWithChannel() chan *Node {

	out := make(chan *Node)

	go func() {
		node.TraverseWithFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()

	return out
}
