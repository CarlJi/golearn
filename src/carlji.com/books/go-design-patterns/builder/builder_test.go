package builder

import "testing"

func TestManunifactorBuilder_Construct(t *testing.T) {
	car := &Car{}
	m := &ManunifactorBuilder{}
	m.SetBuilder(car)
	m.Construct()

	p := car.GetVehicle()
	if p.Size != 4 {
		t.Errorf("failed, size: %d \n", p.Size)
	}

	if p.Wheel != 4 {
		t.Errorf("failed")
	}

	bus := &Bus{}
	m.SetBuilder(bus)
	m.Construct()
	b := bus.GetVehicle()
	if b.Wheel != 4 {
		t.Errorf("failed")
	}

	if b.Size != 16 {
		t.Fatalf("failed")
	}

}
