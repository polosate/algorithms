package exercises

func findLoopStart(list *singleList) *link {
	slowIter := list.GetIterator()
	fastIter := list.GetIterator()

	for fastIter.GetCurrent() != nil && fastIter.GetCurrent().next != nil {
		slowIter.NextLink()
		fastIter.NextLink()
		fastIter.NextLink()
		if slowIter.GetCurrent() == fastIter.GetCurrent() {
			break
		}
	}
	if fastIter.GetCurrent() == nil || fastIter.GetCurrent().next == nil {
		return nil
	}
	slowIter.Reset()
	for fastIter.GetCurrent() != slowIter.GetCurrent() {
		fastIter.NextLink()
		slowIter.NextLink()
	}
	return fastIter.GetCurrent()
}
