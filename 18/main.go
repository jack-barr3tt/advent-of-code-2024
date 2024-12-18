package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jack-barr3tt/gostuff/graphs"
	"github.com/jack-barr3tt/gostuff/maze"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

func decodeNode(name string) types.Point {
	return types.PointFromSlice(stringstuff.GetNums(name))
}

func encodeNode(p types.Point) string {
	return fmt.Sprintf("%d,%d", p[0], p[1])
}

func gridAtTime(bytes []types.Point, size, time int) maze.Maze {
	grid := maze.NewBlankMaze(size, size)
	for i := 0; i < time; i++ {
		grid.Set(bytes[i], '#')
	}
	return grid
}

func graphForGrid(grid maze.Maze, source string) graphs.Graph {
	return graphs.NewVirtualGraph(func(n *graphs.Node) []graphs.Edge {
		pos := decodeNode(n.Name)
		edges := []graphs.Edge{}
		for i := 0; i < 360; i += 90 {
			if next, ok := grid.Move(pos, types.North.Rotate(i)); ok && grid.At(next) != '#' {
				edges = append(edges, graphs.Edge{
					Node: encodeNode(next),
					Cost: 1,
				})
			}
		}
		return edges
	}, source)
}

func main() {
	data, _ := os.ReadFile("input.txt")
	w := 71
	simCount := 1024

	bytes := slicestuff.Map(func(v string) types.Point {
		p := types.PointFromSlice(stringstuff.GetNums(v))
		p[1] = w - 1 - p[1]
		return p
	}, strings.Split(string(data), "\n"))

	startPos := types.Point{0, w - 1}
	endPos := types.Point{w - 1, 0}

	_, length := graphForGrid(gridAtTime(bytes, w, simCount), encodeNode(startPos)).ShortestPath(encodeNode(startPos), encodeNode(endPos), func(n graphs.Node) int { return 1 })

	println("part 1:", length)

	l := 0
	u := len(bytes)
	m := 0
	for l < u {
		m = (l + u) / 2

		_, length := graphForGrid(gridAtTime(bytes, w, m), encodeNode(startPos)).ShortestPath(encodeNode(startPos), encodeNode(endPos), func(n graphs.Node) int { return 1 })

		if length == -1 {
			u = m
		} else {
			l = m + 1
		}
	}

	p := bytes[m]
	p[1] = w - 1 - p[1]

	fmt.Println("part 2:", p)
}
