package exercises

// 1 2 3 5 8 4
func Remove(el *link) {
	if el.next == nil {
		el = nil
	} else {
		el.value = el.next.value
		el.next = el.next.next
	}
}
