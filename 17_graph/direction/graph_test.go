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
	g.addEdge(0, 3)
	g.addEdge(0, 4)
	g.addEdge(1, 4)
	g.addEdge(2, 5)
	g.addEdge(3, 6)
	g.addEdge(4, 6)
	g.addEdge(5, 7)
	g.addEdge(6, 7)

	g.topo()
}
