package builder

type BusAssembler struct {
	v Vehicle
}

func (b *BusAssembler) SetWheels() VehicleAssembler {
	b.v.Wheels = 4 * 2
	return b
}

func (b *BusAssembler) SetSeats() VehicleAssembler {
	b.v.Seats = 30
	return b
}

func (b *BusAssembler) SetStructure() VehicleAssembler {
	b.v.Structure = "Bus"
	return b
}

func (b *BusAssembler) Assemble() Vehicle {
	return b.v
}
