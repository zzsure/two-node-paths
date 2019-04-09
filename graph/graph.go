package graph

import (
"errors"
"github.com/golang-collections/collections/stack"
)

type NodePath struct {
	Distance int64
	Nodes []*Node
}

type Node struct {
	ID int
	edges map[int]int64
}

func (n *Node) AddEdge(e int, d int64) {
	n.edges[e] = d
}

type GraphPaths struct {
	nodes []*Node
	paths []NodePath
	nodeStack *stack.Stack
}

func NewGraphPaths() *GraphPaths {
	g := &GraphPaths{
		nodes: make([]*Node, 1000),
		paths: make([]NodePath, 0),
		nodeStack: stack.New(),
	}
	return g
}

func (g *GraphPaths) AddNode(ID int) *Node {
	for _, v := range g.nodes {
		if v != nil && v.ID == ID {
			return g.nodes[ID]
		}
	}
	n := Node{
		ID: ID,
		edges: make(map[int]int64),
	}
	g.nodes[ID] = &n
	return &n
}

func (g *GraphPaths) AddEdge(s, e int, d int64) error {
	if len(g.nodes) <= s || len(g.nodes) <= e {
		return errors.New("source/distination not found")
	}
	g.nodes[s].AddEdge(e, d)
	return nil
}

func (g *GraphPaths) GetAllPaths(s, e int) ([]NodePath, error) {
	g.getPaths(g.nodes[s], nil, g.nodes[s], g.nodes[e])
	return g.paths, nil
}

func (g *GraphPaths) recordPath() {
	np := NodePath{
		Distance: 0,
		Nodes: make([]*Node, 0),
	}
	tmpStack := stack.New()
	for {
		n := g.nodeStack.Pop()
		tmpStack.Push(n)
		if g.nodeStack.Len() == 0 {
			break
		}
	}
	pre := tmpStack.Pop().(*Node)
	g.nodeStack.Push(pre)
	np.Nodes = append(np.Nodes, pre)
	for ; tmpStack.Len() != 0; {
		curr := tmpStack.Pop().(*Node)
		g.nodeStack.Push(curr)
		np.Nodes = append(np.Nodes, curr)
		np.Distance += pre.edges[curr.ID]
		pre = curr
	}
	g.paths = append(g.paths, np)
}

func (g *GraphPaths) getPaths(cNode, pNode, sNode, eNode *Node) bool {
	if cNode == pNode {
		return false
	}
	if cNode != nil {
		g.nodeStack.Push(cNode)
		if cNode == eNode {
			g.recordPath()
			return true
		}
		for idx, _ := range cNode.edges {
			nNode := g.nodes[idx]
			if nNode == nil {
				break
			}
			if pNode != nil && nNode == sNode || nNode == pNode || g.isNodeInStack(nNode) {
				continue
			}
			if g.getPaths(nNode, cNode, sNode, eNode) {
				g.nodeStack.Pop()
			}
		}
		g.nodeStack.Pop()
		return false
	}
	return false
}

func (g *GraphPaths) isNodeInStack(n *Node) bool {
	isIn := false
	tmpStack := stack.New()
	for {
		ns := g.nodeStack.Pop()
		tmpStack.Push(ns)
		if ns == n {
			isIn = true
		}
		if g.nodeStack.Len() == 0 {
			break
		}
	}
	for {
		tmp := tmpStack.Pop()
		g.nodeStack.Push(tmp)
		if tmpStack.Len() == 0 {
			break
		}
	}
	return isIn
}