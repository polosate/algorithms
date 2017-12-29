package animal_queue

type IAnimal interface {
	GetName() string
}

type animal struct {
	name string
}

func newAnimal(name string) IAnimal {
	return &animal{name: name}
}

func (a *animal) GetName() string {
	return a.name
}

type Dog struct {
	IAnimal
}

func NewDog(name string) Dog {
	dog := newAnimal(name)
	return Dog{dog}
}

type Cat struct {
	IAnimal
}

func NewCat(name string) Cat {
	cat := newAnimal(name)
	return Cat{cat}
}
