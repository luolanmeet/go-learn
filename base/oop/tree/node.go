package tree

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Print() {
	fmt.Printf("%d ", node.Value)
}
