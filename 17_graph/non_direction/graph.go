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
	fmt.Println("======= Depth-first search =======")
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
	fmt.Println()
	for i := 0; i < g.vertexCount; i++ {
		g.vertexList[i].wasVisited = false
	}
}

func (g *graph) bfs() {
	fmt.Println("======= Breadth-first search =======")
	q := newQueue(g.maxVertex)
	g.vertexList[0].wasVisited = true
	g.displayVertex(0)
	q.add(0)
	for !q.isEmpty() {
		v1 := q.remove()
		for {
			v2 := g.getAdjUnvisitedVertex(v1)
			if v2 == -1 {
				break
			}
			g.vertexList[v2].wasVisited = true
			g.displayVertex(v2)
			q.add(v2)
		}
	}
	fmt.Println()
	for i := 0; i < g.vertexCount; i++ {
		g.vertexList[i].wasVisited = false
	}
}

func (g *graph) mstDfs() {
	fmt.Println("======= Minimum spanning tree (depth-first search)=======")
	s := stack.NewIntStack(g.maxVertex)
	g.vertexList[0].wasVisited = true
	s.Push(0)
	for !s.IsEmpty() {
		current, _ := s.Peek()
		v := g.getAdjUnvisitedVertex(int(current))
		if v == -1 {
			s.Pop()
		} else {
			g.vertexList[v].wasVisited = true
			s.Push(int64(v))
			g.displayVertex(int(current))
			g.displayVertex(v)
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for i := 0; i < g.vertexCount; i++ {
		g.vertexList[i].wasVisited = false
	}
}

func (g *graph) mstBfs() {
	fmt.Println("======= Minimum spanning tree (breadth-first search)=======")
	q := newQueue(g.maxVertex)
	g.vertexList[0].wasVisited = true
	q.add(0)
	for !q.IsEmpty() {
		current := q.remove()
		for {
			v := g.getAdjUnvisitedVertex(current)
			if v == -1 {
				break
			}
			g.vertexList[v].wasVisited = true
			q.add(v)
			g.displayVertex(current)
			g.displayVertex(v)
			fmt.Print(" ")
		}
	}
	fmt.Println()
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
	fmt.Print(g.vertexList[v].label)
}
