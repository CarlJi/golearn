package abstract_factory

type LCar struct {
}

func (c *LCar) NumWheels() int {
	return 4
}

func (c *LCar) NumSeats() int {
	return 4
}

func (c *LCar) NumDoors() int {
	return 4
}
