package main

import (
	"os"
	"regexp"

	"github.com/jack-barr3tt/gostuff/parsing"
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

func main() {
	data, _ := os.ReadFile("input.txt")

	part1 := 0
	part2 := 0
	enabled := true

	for source, stmt := parsing.NextToken(string(data), exprs); stmt != ""; source, stmt = parsing.NextToken(source, exprs) {
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
