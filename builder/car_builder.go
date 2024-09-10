package builder

import (
	"github.com/Beretta350/golang-design-patterns/builder/model"
)

// CarBuilder defines the steps for building a car.
type CarBuilder interface {
	SetModel(name string)
	SetEngine(engineType string)
	SetTransmission(transmissionType string)
	SetNumberOfSeats(numSeats int)
	GetCar() model.Car
}
