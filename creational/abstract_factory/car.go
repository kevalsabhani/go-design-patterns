package abstract_factory

type Car interface {
	NumDoors() int
}

type LuxuryCar struct{}

func (c *LuxuryCar) NumDoors() int {
	return 4
}

func (c *LuxuryCar) NumSeats() int {
	return 4
}

func (c *LuxuryCar) NumWheels() int {
	return 4
}

type FamilyCar struct{}

func (*FamilyCar) NumDoors() int {
	return 5
}

func (*FamilyCar) NumWheels() int {
	return 4
}

func (*FamilyCar) NumSeats() int {
	return 5
}
