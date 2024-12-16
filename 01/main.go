package main

import (
	"os"
	"sort"
	"strings"

	numstuff "github.com/jack-barr3tt/gostuff/nums"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

func getLists(data string) ([]int, []int) {
	lines := strings.Split(string(data), "\n")

	left, right := slicestuff.Unzip(slicestuff.Map(func(line string) types.Pair[int, int] {
		return types.PairFromSlice(stringstuff.GetNums(line))
	}, lines))

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

func part1(left, right []int) int {
	combined := slicestuff.Zip(left, right)

	return slicestuff.Reduce(func(pair types.Pair[int, int], acc int) int {
		return acc + numstuff.Abs(pair.Second-pair.First)
	}, combined, 0)
}

func part2(left, right []int) int {
	rightCounts := slicestuff.Frequency(right)

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
