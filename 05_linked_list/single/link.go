package single

import "fmt"

type Link struct {
	iData int
	dData float32
	next  *Link
}

func NewLink(iData int, dData float32) *Link {
	return &Link{
		iData: iData,
		dData: dData,
	}
}

func (l Link) DisplayLink() {
	fmt.Println(l.iData, l.dData)
}
