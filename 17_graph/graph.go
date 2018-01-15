package _7_graph

import "fmt"

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

func (g *graph) displayVertex(v int) {
	fmt.Println(g.vertexList[v].label)
}
