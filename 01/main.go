package main

import (
	"math"
	"os"
	"sort"
	"strings"

	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

func getLists(data string) ([]int, []int) {
	lines := strings.Split(string(data), "\n")

	pairs := slicestuff.Map(func(line string) []int {
		return stringstuff.GetNums(line)
	}, lines)

	left := slicestuff.Map(func(pair []int) int {
		return pair[0]
	}, pairs)
	right := slicestuff.Map(func(pair []int) int {
		return pair[1]
	}, pairs)

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

func part1(left, right []int) int {
	combined := slicestuff.Zip(left, right)

	return slicestuff.Reduce(func(pair types.Pair[int, int], acc int) int {
		return acc + int(math.Abs(float64(pair.Second)-float64(pair.First)))
	}, combined, 0)
}

func freqCounter(list []int) map[int]int {
	counter := make(map[int]int)
	for _, num := range list {
		if _, ok := counter[num]; ok {
			counter[num]++
		} else {
			counter[num] = 1
		}
	}
	return counter
}

func part2(left, right []int) int {
	rightCounts := freqCounter(right)

	return slicestuff.Reduce(func(num int, acc int) int {
		if count, ok := rightCounts[num]; ok {
			return acc + num*count
		}
		return acc
	}, left, 0)
}

func main() {
	data, _ := os.ReadFile("input.txt")

	left, right := getLists(string(data))

	println("part 1:", part1(left, right))
	println("part 2:", part2(left, right))
}
