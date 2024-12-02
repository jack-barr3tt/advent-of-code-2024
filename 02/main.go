package main

import (
	"os"
	"strings"

	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func remove(slice []int, x int) []int {
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:x]...)
	return append(result, slice[x+1:]...)
}

func reportSafe(report []int) bool {
	diffs := []int{}
	for i := 1; i < len(report); i++ {
		diffs = append(diffs, report[i]-report[i-1])
	}

	mul := 1
	if diffs[0] < 0 {
		mul = -1
	}

	return !slicestuff.Some(func(diff int) bool {
		return diff*mul < 1 || diff*mul > 3
	}, diffs)
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	reports := slicestuff.Map(func(line string) []int {
		return stringstuff.GetNums(line)
	}, lines)

	part1 := 0
	part2 := 0

	for _, report := range reports {
		if reportSafe(report) {
			part1++
			part2++
		} else {
			for i := 0; i < len(report); i++ {
				if reportSafe(remove(report, i)) {
					part2++
					break
				}
			}
		}
	}

	println("part 1:", part1)
	println("part 2:", part2)
}
