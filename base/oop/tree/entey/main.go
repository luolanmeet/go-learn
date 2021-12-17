package main

import ()

func main() {

	/**
	 *     0
	 *   1   2
	 * 3       4
	 *        5
	 *      6   7
	 */

	nodes := []Node{
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
}
