package animal_queue

type IAnimalQueue interface {
	Enqueue(a IAnimal)
	DequeueAny() IAnimal
	DequeueCat() Cat
	DequeueDog() Dog
}

type animalQueue struct {
	IAnimalQueue
	dogList ILinkedList
	catList ILinkedList
	order   int
}

func NewAnimalQueue() IAnimalQueue {
	return animalQueue{
		order: -1,
	}
}

func (q *animalQueue) Enqueue(a IAnimal) {
	q.order++
	if animal, ok := a.(Dog); ok {
		q.dogList.addLast(animal, q.order)
		return
	}
	if animal, ok := a.(Cat); ok {
		q.catList.addLast(animal, q.order)
		return
	}
}

func (q *animalQueue) DequeueCat() Cat {
	animal := q.catList.removeFirst()
	cat, _ := animal.(Cat)
	return cat
}

func (q *animalQueue) DequeueDog() Dog {
	animal := q.dogList.removeFirst()
	dog, _ := animal.(Dog)
	return dog
}

func (q *animalQueue) DequeueAny() IAnimal {
	cat := q.catList.peekFirst()
	dog := q.dogList.peekFirst()
	if cat.isOlderThan(dog) {
		return q.catList.removeFirst()
	} else {
		return q.dogList.removeFirst()
	}
}
