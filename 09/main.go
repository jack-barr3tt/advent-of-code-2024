package main

import (
	"math"
	"os"
	"strings"

	slicestuff "github.com/jack-barr3tt/gostuff/slices"
)

func checksumV1(input string) int {
	diskMap := slicestuff.StrsToInts(strings.Split(input, ""))

	s := 0
	e := len(diskMap) - 1
	exS := 0
	checksum := 0

	for s <= e {
		if s%2 == 0 {
			checksum += (s / 2) * (diskMap[s] * (2*exS + diskMap[s] - 1) / 2)
			exS += diskMap[s]
		} else {
			free := diskMap[s]
			for free > 0 {
				remove := int(math.Min(float64(diskMap[e]), float64(free)))
				checksum += (e / 2) * (remove * (2*exS + remove - 1) / 2)
				free -= remove
				diskMap[e] -= remove
				exS += remove
				if diskMap[e] == 0 {
					e -= 2
				}
			}
		}
		s++
	}

	return checksum
}

type Block struct {
	size int
	val  int
}

func checksumV2(input string) int {
	diskMap := slicestuff.StrsToInts(strings.Split(input, ""))

	expMap := []Block{}

	for i, v := range diskMap {
		if i%2 == 0 {
			expMap = append(expMap, Block{size: v, val: i / 2})
		} else {
			expMap = append(expMap, Block{size: v, val: -1})
		}
	}

	for i := len(expMap) - 1; i >= 0; i-- {
		if expMap[i].val == -1 {
			continue
		}

		spaceLoc := slicestuff.FindIndex(func(b Block) bool {
			if b.val == -1 && b.size >= expMap[i].size {
				return true
			}
			return false
		}, expMap)

		if spaceLoc == -1 || spaceLoc >= i {
			continue
		}

		insert := []Block{expMap[i]}
		if expMap[i].size < expMap[spaceLoc].size {
			insert = append(insert, Block{size: expMap[spaceLoc].size - expMap[i].size, val: -1})
		}

		expMap[i] = Block{size: expMap[i].size, val: -1}
		expMap = append(expMap[:spaceLoc], append(insert, expMap[spaceLoc+1:]...)...)
		expMap = joinSpaces(expMap)
	}

	checksum := 0
	i := 0
	for _, b := range expMap {
		if b.val != -1 {
			checksum += b.val * (((i + b.size - i) * (i + i + b.size - 1)) / 2)
		}
		i += b.size
	}

	return checksum
}

func joinSpaces(expMap []Block) []Block {
	for i := 0; i < len(expMap)-1; i++ {
		if expMap[i].val == -1 && expMap[i+1].val == -1 {
			expMap[i] = Block{size: expMap[i].size + expMap[i+1].size, val: -1}
			expMap = append(expMap[:i+1], expMap[i+2:]...)
		}
	}
	return expMap
}

func main() {
	data, _ := os.ReadFile("input.txt")

	println("part 1:", checksumV1(string(data)))
	println("part 2:", checksumV2(string(data)))
}
