package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
	Build(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {

}

const (
	LCarType = 1
	TCarType = 2
)

type CarFactory struct {
}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LCarType:
		return nil, errors.New("LCarType")
	case TCarType:
		return nil, errors.New("TCarType")
	default:
		return nil, errors.New("Unknown CarType")
	}
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorBikeFactory struct {
}

func (m *MotorBikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return nil, nil
	case CruiseMotorbikeType:
		return nil, nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
