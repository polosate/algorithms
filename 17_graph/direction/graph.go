package non_direction

import (
	"fmt"

	"algorithms/03_stack"
)

type graph struct {
	maxVertex   int
	vertexList  []vertex
	adjMat      [][]int
	vertexCount int
	sortedList  []string
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
		sortedList:  make([]string, maxVertex),
	}
}

func (g *graph) addVertex(label string) {
	g.vertexList[g.vertexCount] = newVertex(label)
	g.vertexCount++
}

func (g *graph) addEdge(start, end int) {
	g.adjMat[start][end] = 1
}

func (g *graph) topo() {

}

func (g *graph) noSuccessors() int {

}

func (g *graph) displayVertex(v int) {
	fmt.Print(g.vertexList[v].label)
}
