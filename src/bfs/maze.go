package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := 0; i < row; i++ {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {

	row, col := len(maze), len(maze[0])

	steps := make([][]int, row)
	for i := range steps {
		steps[i] = make([]int, col)
	}

	// 队列中，维护起点
	Q := []point{start}

	for len(Q) > 0 {

		cur := Q[0]
		Q = Q[1:]

		// 如果这样写，会死循环
		// Q := Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// 是否可走
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// 是否走过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			// 是否是起点
			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {

	maze := readMaze("maze.in")
	// for _, row := range maze {
	// 	for _, val := range row {
	// 		fmt.Printf("%d ", val)
	// 	}
	// 	fmt.Println()
	// }

	steps := walk(
		maze,
		point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}

}
