package _5_hash_table

import "fmt"

type dataItem struct {
	key int64
}

func newDataItem(key int64) *dataItem {
	return &dataItem{key: key}
}

func (d dataItem) getKey() int64 {
	return d.key
}

func (d dataItem) displayDataItem() {
	fmt.Print(d.key)
}
