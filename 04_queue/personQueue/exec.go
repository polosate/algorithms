package personQueue

func Exec() {
	shop := NewShop(3)

	for i := 0; i < 30; i++ {
		shop.balancer(*NewPerson(i, 10))
		shop.tick()
	}

	for !shop.Done() {
		shop.tick()
	}
}
