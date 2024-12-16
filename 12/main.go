package main

import (
	"os"

	"github.com/jack-barr3tt/gostuff/maze"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/types"
)

var dirs = []types.Direction{
	types.North,
	types.East,
	types.South,
	types.West,
}
var odirs = []types.Direction{
	types.NorthEast,
	types.SouthEast,
	types.SouthWest,
	types.NorthWest,
}

func findRegion(grid maze.Maze, visited map[types.Point]bool, point types.Point, a rune) (int, int, int) {
	if grid.At(point) != a {
		return 0, 1, 0
	}
	if visited[point] {
		return 0, 0, 0
	}
	visited[point] = true

	na, np, nf := 1, 0, 0

	check := slicestuff.Map(func(dir types.Direction) bool {
		p, ok := grid.Move(point, dir)
		if !ok {
			np += 1
		}
		ta, tp, tf := findRegion(grid, visited, p, a)
		na += ta
		np += tp
		nf += tf
		return !ok || grid.At(p) != a
	}, dirs)

	// external vertices
	for i := 0; i < 4; i++ {
		if check[i] && check[(i+1)%4] {
			nf++
		}
	}

	// internal vertices
	for _, dir := range odirs {
		if p, ok := grid.Move(point, dir); ok && grid.At(p) != a {
			p1, _ := grid.Move(p, types.Direction{dir[0], 0}.Inverse())
			p2, _ := grid.Move(p, types.Direction{0, dir[1]}.Inverse())
			if grid.At(p1) == a && grid.At(p2) == a {
				nf++
			}
		}
	}

	return na, np, nf
}

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := maze.NewMaze(string(data))

	visited := make(map[types.Point]bool)

	part1 := 0
	part2 := 0

	for y := len(grid) - 1; y >= 0; y-- {
		for x := range grid[y] {
			point := types.Point{x, y}
			a, p, f := findRegion(grid, visited, point, grid.At(point))
			part1 += a * p
			part2 += a * f
		}
	}

	println("part 1:", part1)
	println("part 2:", part2)
}
