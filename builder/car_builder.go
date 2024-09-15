package builder

import (
	"github.com/Beretta350/golang-design-patterns/builder/model"
)

// CarBuilder defines the steps for building a car.
type CarBuilder interface {
	SetModel(name string) CarBuilder
	SetEngine(engineType string) CarBuilder
	SetTransmission(transmissionType string) CarBuilder
	SetNumberOfSeats(numSeats int) CarBuilder
	BuildCar() model.Car
}

type CarManager struct {
	builder CarBuilder
}

func NewCarManager() *CarManager {
	return &CarManager{builder: nil}
}

func (c *CarManager) SetBuilder(b CarBuilder) CarBuilder {
	c.builder = b
	return c.builder
}
