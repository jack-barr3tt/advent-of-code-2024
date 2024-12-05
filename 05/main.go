package main

import (
	"os"
	"regexp"
	"sort"

	"github.com/jack-barr3tt/gostuff/parsing"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

var order = regexp.MustCompile(`\d+\|\d+`)
var produce = regexp.MustCompile(`(\d+,)*\d+`)
var toks = []regexp.Regexp{*order, *produce}

func main() {
	data, _ := os.ReadFile("input.txt")

	orderMap := make(map[int][]int)

	part1 := 0
	part2 := 0

	for puzzle, token := parsing.NextToken(string(data), toks); token != ""; puzzle, token = parsing.NextToken(puzzle, toks) {
		nums := stringstuff.GetNums(token)
		if order.MatchString(token) {
			if _, ok := orderMap[nums[0]]; !ok {
				orderMap[nums[0]] = []int{nums[1]}
			} else {
				orderMap[nums[0]] = append(orderMap[nums[0]], nums[1])
			}
		} else if produce.MatchString(token) {
			correct := true

			for i, num := range nums {
				if _, ok := orderMap[num]; ok {
					indexes := slicestuff.Map(func(n int) int {
						return slicestuff.FindIndex(func(x int) bool {
							return x == n
						}, nums)
					}, orderMap[num])

					if slicestuff.Some(func(x int) bool {
						return x != -1 && x < i
					}, indexes) {
						correct = false
					}
				}
			}

			if correct {
				part1 += nums[len(nums)/2]
			} else {
				sort.Slice(nums, func(i, j int) bool {
					if rules, ok := orderMap[nums[i]]; ok {
						if slicestuff.Some(func(x int) bool {
							return x == nums[j]
						}, rules) {
							return true
						}
					}
					return false
				})
				part2 += nums[len(nums)/2]
			}
		}
	}

	println("part 1:", part1)
	println("part 2:", part2)
}
