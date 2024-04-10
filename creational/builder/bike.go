package builder

type BikeAssembler struct {
	v Vehicle
}

func (b *BikeAssembler) SetWheels() VehicleAssembler {
	b.v.Wheels = 2
	return b
}

func (b *BikeAssembler) SetSeats() VehicleAssembler {
	b.v.Seats = 1
	return b
}

func (b *BikeAssembler) SetStructure() VehicleAssembler {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeAssembler) Assemble() Vehicle {
	return b.v
}
