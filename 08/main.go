package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/jack-barr3tt/gostuff/maze"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/types"
)

func getAntennas(data string) []rune {
	antennaRegex := regexp.MustCompile(`[a-zA-Z0-9]`)

	antennaCounts := slicestuff.Frequency(antennaRegex.FindAllString(string(data), -1))

	antennas := []rune{}

	for a, c := range antennaCounts {
		if c >= 2 {
			antennas = append(antennas, rune(a[0]))
		}
	}
	return antennas
}

func getAntinodes(grid maze.Maze, source types.Point, dir types.Direction) ([]string, []string) {
	antinodes := []string{}
	antinodes2 := []string{}

	i := 0
	for {
		p, ok := grid.Move(source, dir.Multiply(i))
		key := fmt.Sprintf("%d,%d", p[0], p[1])
		if ok {
			if i == 1 {
				antinodes = append(antinodes, key)
			}
			antinodes2 = append(antinodes2, key)
		} else {
			break
		}
		i++
	}
	return antinodes, antinodes2
}

func main() {
	data, _ := os.ReadFile("input.txt")

	antennas := getAntennas(string(data))
	grid := maze.NewMaze(string(data))

	antinodes := []string{}
	antinodes2 := []string{}

	for _, a := range antennas {
		locations := grid.LocateAll(a)
		pairs := slicestuff.NCombosUnique(locations, 2)

		for _, pair := range pairs {
			dir := pair[0].DirectionTo(pair[1])

			a1, a2 := getAntinodes(grid, pair[1], dir)
			antinodes = append(antinodes, a1...)
			antinodes2 = append(antinodes2, a2...)

			a1, a2 = getAntinodes(grid, pair[0], dir.Inverse())
			antinodes = append(antinodes, a1...)
			antinodes2 = append(antinodes2, a2...)
		}
	}

	println("part 1:", len(slicestuff.Unique(antinodes)))
	println("part 2:", len(slicestuff.Unique(antinodes2)))
}
