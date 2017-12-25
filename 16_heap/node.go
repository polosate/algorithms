package _6_heap

type node struct {
	iData int64
}

func newNode(key int64) *node {
	return &node{iData: key}
}

func (n *node) getKey() int64 {
	return n.iData
}

func (n *node) setKey(newKey int64) {
	n.iData = newKey
}
