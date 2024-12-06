package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	mapstuff "github.com/jack-barr3tt/gostuff/maps"
	numstuff "github.com/jack-barr3tt/gostuff/nums"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/types"
)

func CanUseForLoop(p1, p2 types.Pair[int, int]) bool {
	x := numstuff.Abs(p1.First - p2.First)
	y := numstuff.Abs(p1.Second - p2.Second)

	return x == 1 || y == 1
}

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := strings.Split(string(data), "\n")

	// find the guard
	guardPos := types.Pair[int, int]{First: 0, Second: 0}
	for y, row := range grid {
		i := regexp.MustCompile(`\^|<|>|V`).FindStringIndex(row)
		if i != nil {
			guardPos = types.Pair[int, int]{First: i[0], Second: y}
			break
		}
	}
	initialPos := types.Pair[int, int]{First: guardPos.First, Second: guardPos.Second}

	visited := make(map[string]bool)

	moves := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	mp := 0
	if grid[guardPos.Second][guardPos.First] == '^' {
		mp = 0
	} else if grid[guardPos.Second][guardPos.First] == '>' {
		mp = 1
	} else if grid[guardPos.Second][guardPos.First] == 'V' {
		mp = 2
	} else if grid[guardPos.Second][guardPos.First] == '<' {
		mp = 3
	}

	initialMp := mp

	for {
		key := fmt.Sprintf("%d,%d", guardPos.First, guardPos.Second)
		visited[key] = true

		move := moves[mp]

		if guardPos.First+move[0] >= len(grid[0]) || guardPos.First+move[0] < 0 || guardPos.Second+move[1] >= len(grid) || guardPos.Second+move[1] < 0 {
			break
		}

		for grid[guardPos.Second+move[1]][guardPos.First+move[0]] == '#' {
			mp = (mp + 1) % 4
			move = moves[mp]
		}

		guardPos.First += move[0]
		guardPos.Second += move[1]
	}

	println("part 1:", len(mapstuff.Keys(visited)))

	pos := [][]int{}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '#' {
				continue
			}
			if x == initialPos.First && y == initialPos.Second {
				continue
			}

			pos = append(pos, []int{y, x})
		}
	}

	vals := slicestuff.ParallelMap(func(p []int) bool {
		x := p[1]
		y := p[0]

		mgrid := make([]string, len(grid))
		copy(mgrid, grid)

		mgrid[y] = mgrid[y][:x] + "#" + mgrid[y][x+1:]
		mp := initialMp
		guardPos := types.Pair[int, int]{First: initialPos.First, Second: initialPos.Second}

		moveHistory := []string{}

		for {
			key := fmt.Sprintf("%d,%d", guardPos.First, guardPos.Second)
			moveHistory = append(moveHistory, key)

			move := moves[mp]

			if guardPos.First+move[0] >= len(mgrid[0]) || guardPos.First+move[0] < 0 || guardPos.Second+move[1] >= len(mgrid) || guardPos.Second+move[1] < 0 {
				break
			}

			for mgrid[guardPos.Second+move[1]][guardPos.First+move[0]] == '#' {
				mp = (mp + 1) % 4
				move = moves[mp]
			}

			guardPos.First += move[0]
			guardPos.Second += move[1]

			if slicestuff.HasRepeatingSuffix(moveHistory, 4) {
				return true
			}
		}

		return false
	}, pos, 10)

	println("part 2:", len(slicestuff.Filter(func(v bool) bool { return v }, vals)))
}
