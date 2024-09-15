package manufacturer

import (
	"github.com/Beretta350/golang-design-patterns/builder"
	"github.com/Beretta350/golang-design-patterns/builder/model"
)

// FordBuilder is a concrete builder for Ford cars.
type FordBuilder struct {
	car model.Car
}

func (b *FordBuilder) SetModel(name string) builder.CarBuilder {
	b.car.Model = name
	return b
}

func (b *FordBuilder) SetEngine(engineType string) builder.CarBuilder {
	b.car.Engine = engineType
	return b
}

func (b *FordBuilder) SetTransmission(transmissionType string) builder.CarBuilder {
	b.car.Transmission = transmissionType
	return b
}

func (b *FordBuilder) SetNumberOfSeats(numSeats int) builder.CarBuilder {
	b.car.Seats = numSeats
	return b
}

func (b *FordBuilder) BuildCar() model.Car {
	b.car.Brand = "Ford"
	return b.car
}
