package main

import (
	"fmt"
	"os"

	"github.com/jack-barr3tt/gostuff/graphs"
	"github.com/jack-barr3tt/gostuff/maze"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

func normaliseDeg(d int) int {
	return ((d % 360) + 360) % 360
}

func decodeNode(name string) (types.Point, int) {
	nums := stringstuff.GetNums(name)
	return types.Point{nums[0], nums[1]}, nums[2]
}

func encodeNode(pos types.Point, dir int) string {
	return fmt.Sprintf("%d_%d_%d", pos[0], pos[1], dir)
}

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := maze.NewMaze(string(data))

	startPos := grid.LocateAll('S')[0]
	endPos := grid.LocateAll('E')[0]

	grid.Set(startPos, '.')
	grid.Set(endPos, '.')

	graph := graphs.NewVirtualGraph(func(n *graphs.Node) []graphs.Edge {
		pos, dir := decodeNode(n.Name)
		edges := []graphs.Edge{{
			Node: encodeNode(pos, normaliseDeg(dir+90)),
			Cost: 1000,
		}, {
			Node: encodeNode(pos, normaliseDeg(dir-90)),
			Cost: 1000,
		}}
		if fwdPos, ok := grid.Move(pos, types.North.Rotate(dir)); ok && grid.At(fwdPos) == '.' {
			edges = append(edges, graphs.Edge{
				Node: encodeNode(fwdPos, dir),
				Cost: 1,
			})
		}
		return edges
	}, encodeNode(startPos, 90))

	bestCost := -1
	bestPaths := [][]string{}

	for i := 0; i < 180; i += 90 {
		paths, cost := graph.AllShortestPaths("1_1_90", encodeNode(endPos, i))
		if cost > 0 && (bestCost == -1 || cost < bestCost) {
			bestCost = cost
			bestPaths = paths
		}
	}

	println("part 1:", bestCost)

	positions := map[types.Point]bool{}

	for _, path := range bestPaths {
		for _, node := range path {
			pos, _ := decodeNode(node)
			positions[pos] = true
		}
	}

	println("part 2:", len(positions))
}
