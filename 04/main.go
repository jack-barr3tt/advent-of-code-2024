package main

import (
	"os"

	"github.com/jack-barr3tt/gostuff/maze"
	"github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/types"
)

func countXmas(grid maze.Maze, point types.Point) int {
	count := 0

	for r := 0; r < 8; r++ {
		found := true
		for i := 0; i < 4; i++ {
			pos, ok := grid.Move(point, types.North.Rotate(45*r).Multiply(i))
			if !ok {
				found = false
				break
			}

			if grid.At(pos) != rune("XMAS"[i]) {
				found = false
				break
			}
		}

		if found {
			count++
		}
	}

	return count
}

func countMas(grid maze.Maze, origin types.Point) int {
	masMaze := maze.NewMaze(`M.M
.A.
S.S`)

	for r := 0; r < 360; r += 90 {
		if grid.SubMazeAt(masMaze.Rotate(r), origin, []rune{'.'}) {
			return 1
		}
	}

	return 0
}

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := maze.NewMaze(string(data))

	xs := grid.LocateAll('X')
	part1 := slices.Reduce(func(c types.Point, a int) int { return a + countXmas(grid, c) }, xs, 0)
	println("part 1:", part1)

	as := grid.LocateAll('A')
	part2 := slices.Reduce(func(c types.Point, a int) int {
		if origin, ok := grid.Move(c, types.SouthWest); ok {
			return a + countMas(grid, origin)
		}
		return a
	}, as, 0)
	println("part 2:", part2)
}
