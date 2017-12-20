package _5_hash_table

import "fmt"

type DataItem struct {
	key int64
}

func NewDataItem(key int64) *DataItem {
	return &DataItem{key: key}
}

func (d DataItem) GetKey() int64 {
	return d.key
}

func (d DataItem) DisplayDataItem() {
	fmt.Print(d.key)
}
