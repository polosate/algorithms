package exercises

func RemoveDups(list *singleList) {

	var inner, outer, next *link
	outer = list.first

	for outer != nil {

		if outer.next == nil {
			return
		}

		inner = outer

		for inner != nil {
			if inner.next == nil {
				break
			}
			next = inner.next
			if next.value == outer.value {
				inner.next = next.next
			}

			inner = inner.next

		}
		outer = outer.next
	}
	return
}
