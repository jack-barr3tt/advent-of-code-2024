package main

import (
	"os"
	"regexp"

	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

var mul = regexp.MustCompile(`mul\(\d+,\d+\)`)
var do = regexp.MustCompile(`do\(\)`)
var dont = regexp.MustCompile(`don't\(\)`)
var exprs = []regexp.Regexp{*mul, *do, *dont}

func execMul(stmt string) int {
	nums := stringstuff.GetNums(stmt)
	return nums[0] * nums[1]
}

func nextToken(s string) (string, string) {
	loc := []int{-1}
	res := ""

	for _, expr := range exprs {
		if match := expr.FindString(s); match != "" {
			if loc[0] == -1 || expr.FindStringIndex(s)[0] < loc[0] {
				loc = expr.FindStringIndex(s)
				res = match
			}
		}
	}

	if loc[0] == -1 {
		return s, res
	}

	return s[loc[1]:], res
}

func main() {
	data, _ := os.ReadFile("input.txt")

	source := string(data)

	part1 := 0
	part2 := 0
	enabled := true
	stmt := ""

	for source, stmt = nextToken(source); stmt != ""; source, stmt = nextToken(source) {
		if mul.MatchString(stmt) {
			res := execMul(stmt)

			part1 += res

			if enabled {
				part2 += res
			}
		}

		if do.MatchString(stmt) {
			enabled = true
		}

		if dont.MatchString(stmt) {
			enabled = false
		}
	}

	println("part 1:", part1)
	println("part 2:", part2)
}
