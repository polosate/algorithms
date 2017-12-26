package exercises

func deleteDups01(list *singleList) {
	dups := map[int64]bool{}

	if list.first.next == nil {
		return
	}
	var prev *link
	cur := list.first
	for cur != nil {
		if dups[cur.value] {
			prev.next = cur.next
		} else {
			dups[cur.value] = true
			prev = cur
		}
		cur = cur.next
	}
}

func deleteDups02(list *singleList) {
	if list.first == nil {
		return
	}

	var runner, current *link
	current = list.first

	for current != nil {
		runner = current
		for runner.next != nil {
			if runner.next.value == current.value {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		current = current.next
	}
	return
}
