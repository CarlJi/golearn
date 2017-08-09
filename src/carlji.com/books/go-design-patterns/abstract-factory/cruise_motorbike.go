package abstract_factory

type CruiseMotorbike struct {
}

func (s *CruiseMotorbike) NumWheels() int {
	return 7
}

func (s *CruiseMotorbike) NumSeats() int {
	return 0
}

func (s *CruiseMotorbike) GetMotorBikeType() int {
	return CruiseMotorbikeType
}
