package abstract_factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMotorbikeFactory(t *testing.T) {
	mFactory, err := BuildFactory(MotorbikeFactoryType)
	assert.Nil(t, err)
	assert.NotNil(t, mFactory)

	// Sport motor bike
	bike, err := mFactory.NewVehicle(SportMotorbikeType)
	assert.Nil(t, err)
	assert.NotNil(t, bike)

	assert.Equal(t, 2, bike.NumWheels())
	assert.Equal(t, 1, bike.NumSeats())
	assert.Equal(t, 1, bike.(Motorbike).GetMotorbikeType())

	// cruise motor bike
	bike, err = mFactory.NewVehicle(CruiseMotorbikeType)
	assert.Nil(t, err)
	assert.NotNil(t, bike)

	assert.Equal(t, 2, bike.NumWheels())
	assert.Equal(t, 2, bike.NumSeats())
	assert.Equal(t, 2, bike.(Motorbike).GetMotorbikeType())
}

func TestCarFactory(t *testing.T) {
	cFactory, err := BuildFactory(CarFactoryType)
	assert.Nil(t, err)
	assert.NotNil(t, cFactory)

	// Luxury car
	car, err := cFactory.NewVehicle(LuxuryCarType)
	assert.Nil(t, err)
	assert.NotNil(t, car)

	assert.Equal(t, 4, car.NumWheels())
	assert.Equal(t, 4, car.NumSeats())
	assert.Equal(t, 4, car.(Car).NumDoors())

	// Family car
	car, err = cFactory.NewVehicle(FamilyCarType)
	assert.Nil(t, err)
	assert.NotNil(t, car)

	assert.Equal(t, 4, car.NumWheels())
	assert.Equal(t, 5, car.NumSeats())
	assert.Equal(t, 5, car.(Car).NumDoors())
}
