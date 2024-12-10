package main

import (
	"fmt"
	"os"

	mapstuff "github.com/jack-barr3tt/gostuff/maps"
	"github.com/jack-barr3tt/gostuff/maze"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func pointKey(point maze.Point) string {
	return fmt.Sprintf("%d,%d", point[0], point[1])
}

func dfs(grid maze.Maze, start maze.Point, visited map[string]bool) []maze.Point {
	if _, ok := visited[pointKey(start)]; ok {
		return []maze.Point{}
	}
	visited[pointKey(start)] = true

	if grid.At(start) == '9' {
		return []maze.Point{start}
	}

	dir := maze.North
	result := []maze.Point{}

	for i := 0; i < 4; i++ {
		if i > 0 {
			dir = dir.RotateDirection(maze.C90)
		}

		newPos, ok := grid.Move(start, dir)

		if ok && stringstuff.GetNum(string(grid.At(newPos))) == stringstuff.GetNum(string(grid.At(start)))+1 {
			// have to copy the visited map because apparently golang maps are reference types
			newVisited := make(map[string]bool)
			for k, v := range visited {
				newVisited[k] = v
			}
			result = append(result, dfs(grid, newPos, newVisited)...)
		}
	}

	return result
}

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := maze.NewMaze(string(data))

	heads := grid.LocateAll('0')

	scores := slicestuff.Map(func(head maze.Point) [2]int {
		heads := dfs(grid, head, make(map[string]bool))
		seen := make(map[string]bool)

		for _, point := range heads {
			seen[pointKey(point)] = true
		}

		return [2]int{len(mapstuff.Keys(seen)), len(heads)}
	}, heads)

	println("part 1:", slicestuff.Reduce(func(a [2]int, b int) int { return a[0] + b }, scores, 0))
	println("part 2:", slicestuff.Reduce(func(a [2]int, b int) int { return a[1] + b }, scores, 0))
}
