package main

import (
	"os"
	"strings"

	"github.com/jack-barr3tt/gostuff/maze"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

type Robot struct {
	pos maze.Point
	dir maze.Direction
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	xlimit, ylimit := 101, 103

	robots := slicestuff.Map(func(v string) Robot {
		robot := stringstuff.GetNums(v)
		pos := maze.Point{robot[0], ylimit - 1 - robot[1]}
		dir := maze.Direction{robot[2], 0 - robot[3]}
		return Robot{pos, dir}
	}, lines)

	part1, part2 := 0, 0
	lsf := 1 << 31
  
	for i := 0; i < xlimit*ylimit; i++ {
		quadrants := []int{0, 0, 0, 0}

		for _, robot := range robots {
			newPos := robot.pos.UnsafeMove(robot.dir.Multiply(i))
			gridPos := maze.Point{((newPos[0] % xlimit) + xlimit) % xlimit, ((newPos[1] % ylimit) + ylimit) % ylimit}

			if gridPos[0] == xlimit/2 || gridPos[1] == ylimit/2 {
				continue
			}

			i := 0
			if gridPos[1] < ylimit/2 {
				i = 2
			}
			if gridPos[0] > xlimit/2 {
				i++
			}
			quadrants[i]++
		}

		sf := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]

		if sf < lsf {
			lsf = sf
			part2 = i
		}

		if i == 100 {
			part1 = sf
		}
	}

	println("part 1:", part1)
  println("part 2:", part2)
}
