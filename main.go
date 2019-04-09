package main

import (
	"demo/graph"
	"fmt"
)

func main() {
	g := graph.NewGraphPaths()
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)
	g.AddEdge(0, 1, 20)
	g.AddEdge(1, 2, 10)
	g.AddEdge(0, 2, 40)
	paths, err := g.GetAllPaths(0, 2)
	if err != nil {
		fmt.Println(err)
	}
	for idx, v := range paths {
		fmt.Printf("distance:%d, path%d:\n", v.Distance, idx)
		for nidx, n := range v.Nodes {
			fmt.Print(n.ID)
			if nidx != len(v.Nodes) - 1 {
				fmt.Print("->")
			} else {
				fmt.Print("\n\n")
			}
		}
	}
}