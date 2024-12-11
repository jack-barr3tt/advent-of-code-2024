package main

import (
	"os"
	"strconv"

	"github.com/jack-barr3tt/gostuff/maps"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func digitCount(num int) int {
	return len(strconv.Itoa(num))
}

func splitNum(num int) []int {
	str := strconv.Itoa(num)
	return stringstuff.GetNums(str[:len(str)/2] + " " + str[len(str)/2:])
}

func pass(nums map[int]int) map[int]int {
	result := make(map[int]int)

	for k, v := range nums {
		if k == 0 {
			result[1] += nums[0]
		} else if digitCount(k)%2 == 0 {
			split := splitNum(k)
			result[split[0]] += v
			result[split[1]] += v
		} else {
			result[k*2024] += v
		}
	}

	return result
}

func main() {
	data, _ := os.ReadFile("input.txt")

	nums := stringstuff.GetNums(string(data))
	numMap1 := slicestuff.Frequency(nums)
	numMap2 := slicestuff.Frequency(nums)

	for i := 0; i < 25; i++ {
		numMap1 = pass(numMap1)
	}
	for i := 0; i < 75; i++ {
		numMap2 = pass(numMap2)
	}

	println("part 1:", slicestuff.Reduce(func(a, b int) int { return a + b }, maps.Values(numMap1), 0))
	println("part 2:", slicestuff.Reduce(func(a, b int) int { return a + b }, maps.Values(numMap2), 0))
}
