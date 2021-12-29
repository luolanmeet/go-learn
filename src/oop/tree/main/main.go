package main

import (
	"fmt"
	"oop/tree"
)

func main() {

	/**
	 *     0
	 *   1   2
	 * 3       4
	 *        5
	 *      6   7
	 */
	nodes := []tree.Node{
		{Value: 0}, {Value: 1}, {Value: 2}, {Value: 3},
		{Value: 4}, {Value: 5}, {Value: 6}, {Value: 7},
	}

	nodes[0].Left = &nodes[1]
	nodes[0].Right = &nodes[2]
	nodes[1].Left = &nodes[3]
	nodes[2].Right = &nodes[4]
	nodes[4].Left = &nodes[5]
	nodes[5].Left = &nodes[6]
	nodes[5].Right = &nodes[7]

	nodes[0].MidRecursion()

	fmt.Println()

	nodes[0].TraverseWithFunc(
		func(node *tree.Node) {
			fmt.Printf("%d ", node.Value)
		},
	)

	fmt.Println()

	c := nodes[0].TraverseWithChannel()
	maxValue := 0
	for node := range c {
		if node.Value > maxValue {
			maxValue = node.Value
		}
	}
	fmt.Println("maxValue is ", maxValue)
}
