package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

//读取迷宫数据
func readMaze(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var row, col int
	//读取并写入
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	//fmt.Println(maze)
	return maze
}

//上下左右四个方向
var directions = [4]point{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (p point) add(d point) point {
	return point{p.i + d.i, p.j + d.j}
}

//获取探索点的值及是否可通行
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []point{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			break
		}

		for _, dir := range directions {
			next := cur.add(dir)
			mVale, ok := next.at(maze)
			if !ok || mVale == 1 {
				//说明有墙
				continue
			}
			sValue, ok := next.at(steps)
			if !ok || sValue != 0 {
				//说明走过了
				continue
			}

			if next == start {
				continue
			}

			curValue, _ := cur.at(steps)
			steps[next.i][next.j] = curValue + 1

			queue = append(queue, next)
		}

	}

	return steps

}

func main() {
	path := "maze/maze.in"
	mazeData := readMaze(path)
	fmt.Println(mazeData)
	steps := walk(mazeData, point{0, 0}, point{len(mazeData) - 1, len(mazeData[0]) - 1})
	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}
}
