package builder

type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

type VehicleAssembler interface {
	SetWheels() VehicleAssembler
	SetSeats() VehicleAssembler
	SetStructure() VehicleAssembler
	Assemble() Vehicle
}
