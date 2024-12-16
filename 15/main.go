package main

import (
	"os"
	"strings"

	"github.com/jack-barr3tt/gostuff/maze"
	"github.com/jack-barr3tt/gostuff/types"
)

func pushVertical(grid *maze.Maze, robotPos *types.Point, pos types.Point, move types.Direction) bool {
	if grid.At(pos) == '.' {
		return true
	}
	aPos := pos
	bPos := pos
	if grid.At(aPos) == '[' {
		bPos = aPos.UnsafeMove(types.East)
	} else {
		aPos = aPos.UnsafeMove(types.West)
	}

	newA, _ := grid.Move(aPos, move)
	newB, _ := grid.Move(bPos, move)
	if grid.At(newA) == '#' || grid.At(newB) == '#' {
		return false
	}

	dummyPoint := types.Point{0, 0}
	dummyGrid := grid.Clone()

	if pushVertical(&dummyGrid, &dummyPoint, newA, move) && pushVertical(&dummyGrid, &dummyPoint, newB, move) {
		pushVertical(grid, &dummyPoint, newA, move)
		pushVertical(grid, &dummyPoint, newB, move)
		grid.Set(aPos, '.')
		grid.Set(bPos, '.')
		grid.Set(newA, '[')
		grid.Set(newB, ']')
		robotPos[0] = pos[0]
		robotPos[1] = pos[1]
		return true
	} else {
		return false
	}
}

func main() {
	data, _ := os.ReadFile("input.txt")

	parts := strings.Split(string(data), "\n\n")

	grid := maze.NewMaze(string(parts[0]))

	moves := []types.Direction{}
	for _, m := range parts[1] {
		if m == '^' {
			moves = append(moves, types.North)
		} else if m == 'v' {
			moves = append(moves, types.South)
		} else if m == '>' {
			moves = append(moves, types.East)
		} else if m == '<' {
			moves = append(moves, types.West)
		}
	}

	robotPos := grid.LocateAll('@')[0]

	for _, move := range moves {
		grid.Set(robotPos, '.')
		newPos, _ := grid.Move(robotPos, move)
		if grid.At(newPos) == '.' {
			robotPos = newPos
		} else if grid.At(newPos) == 'O' {
			for i := 1; ; i++ {
				searchPos, _ := grid.Move(newPos, move.Multiply(i))
				if grid.At(searchPos) == '.' {
					grid.Set(newPos, '.')
					grid.Set(searchPos, 'O')
					robotPos = newPos
					break
				} else if grid.At(searchPos) != 'O' {
					break
				}
			}
		}
		grid.Set(robotPos, '@')
	}

	boxes := grid.LocateAll('O')

	part1 := 0

	for _, box := range boxes {
		part1 += box[0] + 100*(len(grid)-1-box[1])
	}

	println("part 1:", part1)

	bigraw := parts[0]
	bigraw = strings.ReplaceAll(bigraw, "#", "##")
	bigraw = strings.ReplaceAll(bigraw, "O", "[]")
	bigraw = strings.ReplaceAll(bigraw, ".", "..")
	bigraw = strings.ReplaceAll(bigraw, "@", "@.")

	grid = maze.NewMaze(bigraw)
	robotPos = grid.LocateAll('@')[0]

	for _, move := range moves {
		grid.Set(robotPos, '.')
		newPos, _ := grid.Move(robotPos, move)
		if grid.At(newPos) == '.' {
			robotPos = newPos
		} else if grid.At(newPos) == '[' || grid.At(newPos) == ']' {
			if move == types.East || move == types.West {
				for i := 1; ; i++ {
					searchPos, _ := grid.Move(newPos, move.Multiply(i))
					if grid.At(searchPos) == '.' {
						for j := i; j > 0; j-- {
							currPos, _ := grid.Move(newPos, move.Multiply(j))
							prevPos, _ := grid.Move(newPos, move.Multiply(j-1))
							grid.Set(currPos, grid.At(prevPos))
							grid.Set(prevPos, '.')
						}
						robotPos = newPos
						break
					} else if grid.At(searchPos) != '[' && grid.At(searchPos) != ']' {
						break
					}
				}
			} else {
				pushVertical(&grid, &robotPos, newPos, move)
			}
		}
		grid.Set(robotPos, '@')
	}

	boxes = grid.LocateAll('[')

	part2 := 0

	for _, box := range boxes {
		part2 += box[0] + 100*(len(grid)-1-box[1])
	}

	println("part 2:", part2)
}
