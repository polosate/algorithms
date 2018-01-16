package non_direction

import "testing"

func TestGraphTraverse(t *testing.T) {
	g := newGraph(10)
	g.addVertex("A") // 0
	g.addVertex("B") // 1
	g.addVertex("C") // 2
	g.addVertex("D") // 3
	g.addVertex("E") // 4
	g.addVertex("F") // 5
	g.addVertex("G") // 6
	g.addVertex("H") // 7
	g.addVertex("I") // 8
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 3)
	g.addEdge(0, 4)
	g.addEdge(1, 5)
	g.addEdge(5, 7)
	g.addEdge(3, 6)
	g.addEdge(6, 8)

	g.dfs()
	g.bfs()
}

func TestGraphMST(t *testing.T) {
	g := newGraph(5)
	g.addVertex("A") // 0
	g.addVertex("B") // 1
	g.addVertex("C") // 2
	g.addVertex("D") // 3
	g.addVertex("E") // 4

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 3)
	g.addEdge(0, 4)
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(2, 3)
	g.addEdge(2, 4)
	g.addEdge(3, 4)

	g.mst()
}
