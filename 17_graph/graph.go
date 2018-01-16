package _7_graph

import (
	"fmt"

	stack "algorithms/03_stack"
)

type graph struct {
	maxVertex   int
	vertexList  []vertex
	adjMat      [][]int
	vertexCount int
}

func newGraph(maxVertex int) graph {
	adjMat := make([][]int, maxVertex)
	for i := 0; i < maxVertex; i++ {
		a := make([]int, maxVertex)
		adjMat[i] = a
	}
	return graph{
		maxVertex:   maxVertex,
		vertexList:  make([]vertex, maxVertex),
		adjMat:      adjMat,
		vertexCount: 0,
	}
}

func (g *graph) addVertex(label string) {
	g.vertexList[g.vertexCount] = newVertex(label)
	g.vertexCount++
}

func (g *graph) addEdge(start, end int) {
	g.adjMat[start][end] = 1
	g.adjMat[end][start] = 1
}

func (g *graph) dfs() {
	s := stack.NewIntStack(g.maxVertex)
	g.vertexList[0].wasVisited = true
	g.displayVertex(0)
	s.Push(0)

	for !s.IsEmpty() {
		a, _ := s.Peek()
		v := g.getAdjUnvisitedVertex(int(a))
		if v == -1 {
			s.Pop()
		} else {
			g.vertexList[v].wasVisited = true
			g.displayVertex(v)
			s.Push(int64(v))
		}
	}

	for i := 0; i < g.vertexCount; i++ {
		g.vertexList[i].wasVisited = false
	}
}

func (g *graph) getAdjUnvisitedVertex(v int) int {
	for i := 0; i < g.vertexCount; i++ {
		if g.adjMat[v][i] == 1 && !g.vertexList[i].wasVisited {
			return i
		}
	}
	return -1
}

func (g *graph) displayVertex(v int) {
	fmt.Println(g.vertexList[v].label)
}
