package main

import (
	"math"
	"os"
	"strings"

	"github.com/jack-barr3tt/gostuff/lines"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func findPresses(ai, bi, pi []int) (int, int) {
	x, y, ok := lines.NewAXBYC(ai[0], bi[0], pi[0]).IntersectsAt(lines.NewAXBYC(ai[1], bi[1], pi[1]))
	if ok && math.Round(x) == x && math.Round(y) == y {
		return int(math.Round(x)), int(math.Round(y))
	}
	return 0, 0
}

func main() {
	data, _ := os.ReadFile("input.txt")

	chunks := strings.Split(string(data), "\n\n")

	part1, part2 := 0, 0

	for _, chunk := range chunks {
		nums := stringstuff.GetNums(chunk)
		a, b, prize := []int{nums[0], nums[1]}, []int{nums[2], nums[3]}, []int{nums[4], nums[5]}

		a1, b1 := findPresses(a, b, prize)
		part1 += a1*3 + b1

		a2, b2 := findPresses(a, b, []int{prize[0] + 10000000000000, prize[1] + 10000000000000})
		part2 += a2*3 + b2
	}

	println("part 1:", part1)
	println("part 2:", part2)
}
