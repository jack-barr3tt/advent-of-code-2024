package main

import (
	"os"
	"regexp"

	"github.com/jack-barr3tt/gostuff/parsing"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func toFloat(a []int) []float64 {
	return []float64{float64(a[0]), float64(a[1])}
}

func findPresses(ai, bi, pi []int) (int, int) {
	a, b, p := toFloat(ai), toFloat(bi), toFloat(pi)
	ca := ((p[0] * b[1]) - (p[1] * b[0])) / ((a[0] * b[1]) - (a[1] * b[0]))
	cb := (p[0] - a[0]*ca) / b[0]
	if ca != float64(int64(ca)) || cb != float64(int64(cb)) {
		return 0, 0
	}
	return int(ca), int(cb)
}

func main() {
	data, _ := os.ReadFile("input.txt")

	buttonRegex := regexp.MustCompile(`Button (A|B): X\+([0-9]+), Y\+([0-9]+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)

	rest, stmt := string(data), ""

	part1, part2 := 0, 0

	for {
		rest, stmt = parsing.NextToken(rest, []regexp.Regexp{*buttonRegex})
		if stmt == "" {
			break
		}
		a := stringstuff.GetNums(stmt)
		rest, stmt = parsing.NextToken(rest, []regexp.Regexp{*buttonRegex})
		b := stringstuff.GetNums(stmt)
		rest, stmt = parsing.NextToken(rest, []regexp.Regexp{*prizeRegex})
		prize := stringstuff.GetNums(stmt)

		a1, b1 := findPresses(a, b, prize)
		part1 += a1*3 + b1

		a2, b2 := findPresses(a, b, []int{prize[0] + 10000000000000, prize[1] + 10000000000000})
		part2 += a2*3 + b2
	}

	println("part 1:", part1)
	println("part 2:", part2)
}
