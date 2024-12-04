package main

import (
	"os"
	"strings"
)

func countXmas(puzzle []string, x, y int) int {
	if puzzle[x][y] != 'X' {
		return 0
	}
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	count := 0

	for _, move := range moves {
		found := true
		for i := 1; i < 4; i++ {
			nx := x + i*move[0]
			ny := y + i*move[1]
			if nx < 0 || nx >= len(puzzle) || ny < 0 || ny >= len(puzzle[0]) {
				found = false
				break
			}

			c := puzzle[nx][ny]
			if (i == 1 && c != 'M') || (i == 2 && c != 'A') || (i == 3 && c != 'S') {
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

func countMas(puzzle []string, x, y int) int {
	if x+2 >= len(puzzle) || y+2 >= len(puzzle[0]) {
		return 0
	}

	count := 0

	forms := [][]string{
		{
			"M.S",
			".A.",
			"M.S",
		},
		{
			"M.M",
			".A.",
			"S.S",
		},
		{
			"S.M",
			".A.",
			"S.M",
		},
		{
			"S.S",
			".A.",
			"M.M",
		},
	}

	checkForm := func(form []string) bool {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if form[i][j] != '.' && puzzle[x+i][y+j] != form[i][j] {
					return false
				}
			}
		}
		return true
	}

	for _, form := range forms {
		if checkForm(form) {
			count++
		}
	}

	return count
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	part1 := 0
	part2 := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			part1 += countXmas(lines, i, j)
			part2 += countMas(lines, i, j)
		}
	}

	println("part 1:", part1)
	println("part 2:", part2)
}
