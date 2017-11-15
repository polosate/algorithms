package exercises

func RemoveDups(list *singleList) {
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
