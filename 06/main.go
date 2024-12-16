package main

import (
	"fmt"
	"os"

	mapstuff "github.com/jack-barr3tt/gostuff/maps"
	"github.com/jack-barr3tt/gostuff/maze"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/types"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := maze.NewMaze(string(data))

	guardAppearances := []rune{'^', '>', 'V', '<'}
	guardPos := slicestuff.FlatMap(func(g rune) []types.Point {
		return grid.LocateAll(g)
	}, guardAppearances)[0]

	initialPos := guardPos.Clone()

	visited := make(map[string]bool)

	var dir types.Direction
	if grid.At(guardPos) == '^' {
		dir = types.North
	} else if grid.At(guardPos) == '>' {
		dir = types.East
	} else if grid.At(guardPos) == 'V' {
		dir = types.South
	} else if grid.At(guardPos) == '<' {
		dir = types.West
	}
	initDir := dir

	for {
		key := fmt.Sprintf("%d,%d", guardPos[0], guardPos[1])
		visited[key] = true

		newPos, ok := grid.Move(guardPos, dir)
		if !ok {
			break
		}

		for ; grid.At(newPos) == '#'; newPos, ok = grid.Move(guardPos, dir) {
			dir = dir.Rotate(90)
		}

		guardPos = newPos
	}

	println("part 1:", len(mapstuff.Keys(visited)))

	pos := []types.Point{}
	for y := range grid {
		for x := range grid[y] {
			if grid.At(types.Point{x, y}) == '#' || x == initialPos[0] && y == initialPos[1] {
				continue
			}

			pos = append(pos, types.Point{x, y})
		}
	}

	vals := slicestuff.ParallelMap(func(p types.Point) bool {
		mgrid := grid.Clone()
		mgrid.Set(p, '#')
		dir := initDir
		guardPos := initialPos.Clone()
		moveHistory := []string{}

		for {
			key := fmt.Sprintf("%d,%d", guardPos[0], guardPos[1])
			moveHistory = append(moveHistory, key)

			newPos, ok := mgrid.Move(guardPos, dir)
			if !ok {
				break
			}

			for ; mgrid.At(newPos) == '#'; newPos, ok = mgrid.Move(guardPos, dir) {
				dir = dir.Rotate(90)
			}

			guardPos = newPos

			if slicestuff.HasRepeatingSuffix(moveHistory, 4) {
				return true
			}
		}

		return false
	}, pos, 10)

	println("part 2:", len(slicestuff.Filter(func(v bool) bool { return v }, vals)))
}
