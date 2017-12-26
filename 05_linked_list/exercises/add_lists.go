package exercises

func addLists(l1, l2 *singleList) *singleList {
	f := addListsRec(l1.first, l2.first, 0)
	res := NewSingleList()
	res.first = f
	return res
}

func addListsRec(l1, l2 *link, carry int64) *link {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	value := carry
	if l1 != nil {
		value += l1.value
	}

	if l2 != nil {
		value += l2.value
	}
	res := NewLink(value % 10)
	if l1 == nil {
		l1 = nil
	} else {
		l1 = l1.next
	}
	if l2 == nil {
		l2 = nil
	} else {
		l2 = l2.next
	}
	if value > 10 {
		carry = 1
	} else {
		carry = 0
	}
	next := addListsRec(l1, l2, carry)
	res.next = next
	return res
}
