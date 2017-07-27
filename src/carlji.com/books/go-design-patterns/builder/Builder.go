package builder

type BuildProcess interface {
	SetWheel() BuildProcess
	SetSize() BuildProcess
}

type ManunifactorBuilder struct {
	Entry BuildProcess
}

func (m *ManunifactorBuilder) SetBuilder(b BuildProcess) {
	m.Entry = b
}

func (m *ManunifactorBuilder) Construct() {
	m.Entry.SetSize().SetWheel()
}

type VehicleProduct struct {
	Wheel int
	Size  int
}

type Car struct {
	v VehicleProduct
}

func (c *Car) SetWheel() BuildProcess {
	c.v.Wheel = 4
	return c
}

func (c *Car) SetSize() BuildProcess {
	c.v.Size = 4
	return c
}

func (c *Car) GetVehicle() VehicleProduct {
	return c.v
}

type Bus struct {
	v VehicleProduct
}

func (b *Bus) SetWheel() BuildProcess {
	b.v.Wheel = 4
	return b
}

func (b *Bus) SetSize() BuildProcess {
	b.v.Size = 16
	return b
}

func (b *Bus) GetVehicle() VehicleProduct {
	return b.v
}
