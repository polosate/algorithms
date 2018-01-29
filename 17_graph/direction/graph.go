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

func (g *graph) dfs() {
	for i := 0; i < g.vertexCount; i++ {
		g.displayVertex(i)
		g.vertexList[i].wasVisited = true
		s := stack.NewIntStack(g.vertexCount)
		s.Push(int64(i))
		for !s.IsEmpty() {
			a, _ := s.Peek()
			v := g.getAdjUnvisitedVerte(int(a))
			if v == -1 {
				s.Pop()
			} else {
				g.vertexList[v].wasVisited = true
				g.displayVertex(v)
				s.Push(int64(v))
			}
		}
		fmt.Println()
		for j := 0; j < g.vertexCount; j++ {
			g.vertexList[j].wasVisited = false
		}
	}
}

func (g *graph) getAdjUnvisitedVerte(i int) int {
	for j := 0; j < g.vertexCount; j++ {
		if g.adjMat[i][j] == 1 && !g.vertexList[j].wasVisited {
			return j
		}
	}
	return -1
}

func (g *graph) topo() {
	originCount := g.vertexCount
	for g.vertexCount > 0 {
		curVertex := g.noSuccessors()
		if curVertex == -1 {
			fmt.Println("ERROR: Graph has loops")
			return
		}
		g.sortedList[g.vertexCount-1] = g.vertexList[curVertex].label
		g.deleteVertex(curVertex)
	}
	fmt.Println("Topologically sorted order:")
	for i := 0; i < originCount; i++ {
		fmt.Print(g.sortedList[i], " ")
	}
	fmt.Println()
}

func (g *graph) noSuccessors() int {
	var isEdge bool
	for row := 0; row < g.vertexCount; row++ {
		isEdge = false
		for col := 0; col < g.vertexCount; col++ {
			if g.adjMat[row][col] == 1 {
				isEdge = true
				break
			}
		}
		if !isEdge {
			return row
		}
	}
	return -1
}

func (g *graph) deleteVertex(delVert int) {
	if delVert != g.vertexCount-1 {
		for i := delVert; i < g.vertexCount-1; i++ {
			g.vertexList[i] = g.vertexList[i+1]
		}
		for row := delVert; row < g.vertexCount-1; row++ {
			g.moveRowUp(row)
		}
		for col := delVert; col < g.vertexCount-1; col++ {
			g.moveColLeft(col)
		}
	}
	g.vertexCount--
}

func (g *graph) moveRowUp(row int) {
	for col := 0; col < g.vertexCount; col++ {
		g.adjMat[row][col] = g.adjMat[row+1][col]
	}
}

func (g *graph) moveColLeft(col int) {
	for row := 0; row < g.vertexCount; row++ {
		g.adjMat[row][col] = g.adjMat[row][col+1]
	}
}

func (g *graph) displayVertex(v int) {
	fmt.Print(g.vertexList[v].label)
}

func (g *graph) displayMatrix() {
	for i := 0; i < g.vertexCount; i++ {
		for j := 0; j < g.vertexCount; j++ {
			fmt.Print(g.adjMat[i][j], "  ")
		}
		fmt.Println()
	}
	fmt.Println("-----------------------")
}

func (g *graph) displayList() {
	for j := 0; j < g.vertexCount; j++ {
		fmt.Print(g.vertexList[j].label, "  ")
	}
	fmt.Println()
	fmt.Println("-----------------------")
}
