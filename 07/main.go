package main

import (
	"os"
	"strconv"
	"strings"

	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func p1Valid(curr, target int, nums []int) bool {
	if len(nums) == 0 {
		return curr == target
	}

	return p1Valid(curr+nums[0], target, nums[1:]) || p1Valid(curr*nums[0], target, nums[1:])
}

func concatNums(a, b int) int {
	return stringstuff.GetNum(strconv.Itoa(a) + strconv.Itoa(b))
}

func p2Valid(curr, target int, nums []int) bool {
	if len(nums) == 0 {
		return curr == target
	}

	return p2Valid(curr+nums[0], target, nums[1:]) || p2Valid(curr*nums[0], target, nums[1:]) || p2Valid(concatNums(curr, nums[0]), target, nums[1:])
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	part1 := 0
	part2 := 0

	for _, line := range lines {
		parts := strings.Split(line, ":")

		target := stringstuff.GetNum(parts[0])
		nums := stringstuff.GetNums(parts[1])

		if p1Valid(nums[0], target, nums[1:]) {
			part1 += target
		}
		if p2Valid(nums[0], target, nums[1:]) {
			part2 += target
		}
	}

	println("part 1:", part1)
	println("part 2:", part2)
}
