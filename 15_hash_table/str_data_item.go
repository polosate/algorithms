package _5_hash_table

import "fmt"

type StrDataItem struct {
	key string
}

func NewStrDataItem(key string) *StrDataItem {
	return &StrDataItem{key: key}
}

func (d StrDataItem) GetKey() string {
	return d.key
}

func (d StrDataItem) DisplayDataItem() {
	fmt.Print(d.key)
}
