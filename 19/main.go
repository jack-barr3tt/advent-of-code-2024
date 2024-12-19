package main

import (
	"os"
	"regexp"
	"strings"

	"github.com/jack-barr3tt/gostuff/maps"
	"github.com/jack-barr3tt/gostuff/nums"
	"github.com/jack-barr3tt/gostuff/slices"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	available := slices.Frequency(regexp.MustCompile(`[a-z]+`).FindAllString(strings.Split(string(data), "\n\n")[0], -1))
	required := strings.Split(strings.Split(string(data), "\n\n")[1], "\n")
	longest := maps.Reduce(func(acc int, k string, _ int) int { return nums.Max(len(k), acc) }, available, 0)

	memo := map[string]int{}

	var dp func(target string) int
	dp = func(target string) int {
		if v, ok := memo[target]; ok {
			return v
		}
		if target == "" {
			return 1
		}
		count := 0
		for i := 1; i <= longest && i <= len(target); i++ {
			if _, ok := available[target[:i]]; ok {
				if r := dp(target[i:]); r > 0 {
					count += r
				}
			}
		}
		memo[target] = count
		return count
	}

	res := slices.Map(func(r string) int { return dp(r) }, required)

	println("part 1:", slices.CountIf(func(v int) bool { return v > 0 }, res))
	println("part 2:", slices.Sum(res))
}
