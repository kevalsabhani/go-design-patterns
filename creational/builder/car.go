package builder

type CarAssembler struct {
	v Vehicle
}

func (c *CarAssembler) SetWheels() VehicleAssembler {
	c.v.Wheels = 4
	return c
}

func (c *CarAssembler) SetSeats() VehicleAssembler {
	c.v.Seats = 4
	return c
}

func (c *CarAssembler) SetStructure() VehicleAssembler {
	c.v.Structure = "Car"
	return c
}

func (c *CarAssembler) Assemble() Vehicle {
	return c.v
}
