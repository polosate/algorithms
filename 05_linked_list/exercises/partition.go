package exercises

func partition(list *singleList, k int64) *singleList {
	if list.IsEmpty() || list.first.next == nil {
		return list
	}
	before := NewSingleList()
	after := NewSingleList()

	cur := list.first
	var next *link
	for cur != nil {
		next = cur.next
		if cur.value < k {
			if before.IsEmpty() {
				before.first = cur
				cur.next = nil
			} else {
				cur.next = before.first
				before.first = cur
			}
		} else {
			if after.IsEmpty() {
				after.first = cur
				cur.next = nil
			} else {
				cur.next = after.first
				after.first = cur
			}
		}
		cur = next
	}

	if before.IsEmpty() {
		return after
	}
	cur = before.first
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = after.first
	return before

}
