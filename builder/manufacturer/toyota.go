package manufacturer

import (
	"github.com/Beretta350/golang-design-patterns/builder"
	"github.com/Beretta350/golang-design-patterns/builder/model"
)

// ToyotaBuilder is a concrete builder for Toyota cars.
type ToyotaBuilder struct {
	car model.Car
}

func (b *ToyotaBuilder) SetModel(name string) builder.CarBuilder {
	b.car.Model = name
	return b
}

func (b *ToyotaBuilder) SetEngine(engineType string) builder.CarBuilder {
	b.car.Engine = engineType
	return b
}

func (b *ToyotaBuilder) SetTransmission(transmissionType string) builder.CarBuilder {
	b.car.Transmission = transmissionType
	return b
}

func (b *ToyotaBuilder) SetNumberOfSeats(numSeats int) builder.CarBuilder {
	b.car.Seats = numSeats
	return b
}

func (b *ToyotaBuilder) BuildCar() model.Car {
	b.car.Brand = "Toyota"
	return b.car
}
