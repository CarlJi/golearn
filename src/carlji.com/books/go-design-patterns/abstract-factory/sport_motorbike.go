package abstract_factory

type SportMotorybike struct {
}

func (s *SportMotorybike) NumWheels() int {
	return 5
}

func (s *SportMotorybike) NumSeats() int {
	return 1
}

func (s *SportMotorybike) GetMotorBikeType() int {
	return SportMotorbikeType
}


