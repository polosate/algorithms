package _4_2_3_4_Tree

import "fmt"

type dataItem struct {
	key int64
}

func newDataItem(key int64) *dataItem {
	return &dataItem{key: key}
}

func (di dataItem) displayDataItem() {
	fmt.Printf("/%d", di.key)
}
