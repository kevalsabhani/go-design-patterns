package builder

type Manufacturer struct {
	v VehicleAssembler
}

func (m *Manufacturer) Construct() Vehicle {
	return m.v.SetWheels().SetSeats().SetStructure().Assemble()
}

func (m *Manufacturer) SetAssembler(v VehicleAssembler) {
	m.v = v
}
