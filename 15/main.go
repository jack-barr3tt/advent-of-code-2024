package main

import (
	"os"
	"strings"

	"github.com/jack-barr3tt/gostuff/maze"
)

func pushVertical(grid *maze.Maze, robotPos *maze.Point, pos maze.Point, move maze.Direction) bool {
	if grid.At(pos) == '.' {
		return true
	}
	aPos := pos
	bPos := pos
	if grid.At(aPos) == '[' {
		bPos = aPos.UnsafeMove(maze.East)
	} else {
		aPos = aPos.UnsafeMove(maze.West)
	}

	newA, _ := grid.Move(aPos, move)
	newB, _ := grid.Move(bPos, move)
	if grid.At(newA) == '#' || grid.At(newB) == '#' {
		return false
	}

	dummyPoint := maze.Point{0, 0}
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

	moves := []maze.Direction{}
	for _, m := range parts[1] {
		if m == '^' {
			moves = append(moves, maze.North)
		} else if m == 'v' {
			moves = append(moves, maze.South)
		} else if m == '>' {
			moves = append(moves, maze.East)
		} else if m == '<' {
			moves = append(moves, maze.West)
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
			if move == maze.East || move == maze.West {
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
