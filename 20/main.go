package main

import (
	"fmt"
	"os"

	"github.com/jack-barr3tt/gostuff/graphs"
	"github.com/jack-barr3tt/gostuff/maze"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

var (
	NotCheated = 0
	Cheating   = 1
	Cheated    = 2
)

func encodeNode(p types.Point, cs int) string {
	return fmt.Sprintf("%d,%d,%d", p[0], p[1], cs)
}

func decodeNode(n string) (types.Point, int) {
	nums := stringstuff.GetNums(n)
	return types.Point{nums[0], nums[1]}, nums[2]
}

func cheatsEdgeGenerator(grid maze.Maze, cantCheatAt map[string]bool) func(n *graphs.Node) []graphs.Edge {
	return func(n *graphs.Node) []graphs.Edge {
		pos, cs := decodeNode(n.Name)
		edges := []graphs.Edge{}
		for i := 0; i < 360; i += 90 {
			if newPos, ok := grid.Move(pos, types.North.Rotate(i)); ok {
				if grid.At(newPos) == '#' {
					// if you haven't cheated yet, and a wall is in the way, you can start cheating
					if cs == NotCheated && !cantCheatAt[encodeNode(newPos, Cheating)] {
						edges = append(edges, graphs.Edge{
							Node: encodeNode(newPos, Cheating),
							Cost: 1,
						})
					}
				} else {
					ncs := cs
					if ncs >= Cheating && ncs < Cheated {
						ncs++
					}
					edges = append(edges, graphs.Edge{
						Node: encodeNode(newPos, ncs),
						Cost: 1,
					})
				}
			}
		}
		return edges
	}
}

func legitEdgeGenerator(grid maze.Maze) func(n *graphs.Node) []graphs.Edge {
	return func(n *graphs.Node) []graphs.Edge {
		pos, _ := decodeNode(n.Name)
		edges := []graphs.Edge{}
		for i := 0; i < 360; i += 90 {
			if newPos, ok := grid.Move(pos, types.North.Rotate(i)); ok && grid.At(newPos) != '#' {
				edges = append(edges, graphs.Edge{
					Node: encodeNode(newPos, NotCheated),
					Cost: 1,
				})
			}
		}
		return edges
	}
}

func main() {
	data, _ := os.ReadFile("input.txt")

	grid := maze.NewMaze(string(data))

	startPos := grid.LocateAll('S')[0]
	endPos := grid.LocateAll('E')[0]

	lg := graphs.NewVirtualGraph(legitEdgeGenerator(grid), encodeNode(startPos, NotCheated))
	_, legitLength := lg.ShortestPath(encodeNode(startPos, NotCheated), encodeNode(endPos, NotCheated), func(n graphs.Node) int { return 1 })

	cantCheatAt := map[string]bool{}
	part1 := 0

	for {
		graph := graphs.NewVirtualGraph(cheatsEdgeGenerator(grid, cantCheatAt), encodeNode(startPos, NotCheated))
		paths, length := graph.AllShortestPaths(encodeNode(startPos, NotCheated), encodeNode(endPos, Cheated), func(n graphs.Node) int { return 1 })
		if legitLength-length < 100 {
			break
		}
		for _, path := range paths {
			for _, node := range path {
				_, cs := decodeNode(node)
				if cs == Cheating {
					cantCheatAt[node] = true
				}
			}
		}
		part1 += len(paths)
	}

	println("part 1:", part1)
}
