package abstract_factory

type TCar struct {
}

func (c *TCar) NumWheels() int {
	return 5
}

func (c *TCar) NumDoors() int {
	return 5
}

func (c *TCar) NumSeats() int {
	return 5
}
