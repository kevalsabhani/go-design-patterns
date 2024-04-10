package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarAssembler(t *testing.T) {
	expectedCar := Vehicle{
		Wheels:    4,
		Seats:     4,
		Structure: "Car",
	}
	cb := &CarAssembler{}
	m := &Manufacturer{}
	m.SetAssembler(cb)
	car := m.Construct()

	assert.Equal(t, expectedCar, car)
}

func TestBusAssembler(t *testing.T) {
	expectedBus := Vehicle{
		Wheels:    4 * 2,
		Seats:     30,
		Structure: "Bus",
	}
	bb := &BusAssembler{}
	m := &Manufacturer{}
	m.SetAssembler(bb)
	bus := m.Construct()

	assert.Equal(t, expectedBus, bus)
}

func TestBikeAssembler(t *testing.T) {
	expectedBus := Vehicle{
		Wheels:    2,
		Seats:     1,
		Structure: "Bike",
	}
	bb := &BikeAssembler{}
	m := &Manufacturer{}
	m.SetAssembler(bb)
	bike := m.Construct()

	assert.Equal(t, expectedBus, bike)
}
