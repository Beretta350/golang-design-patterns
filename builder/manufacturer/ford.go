package manufacturer

import "github.com/Beretta350/golang-design-patterns/builder/model"

// FordBuilder is a concrete builder for Ford cars.
type FordBuilder struct {
	car model.Car
}

func (b *FordBuilder) SetModel(name string) {
	b.car.Model = name
}

func (b *FordBuilder) SetEngine(engineType string) {
	b.car.Engine = engineType
}

func (b *FordBuilder) SetTransmission(transmissionType string) {
	b.car.Transmission = transmissionType
}

func (b *FordBuilder) SetNumberOfSeats(numSeats int) {
	b.car.Seats = numSeats
}

func (b *FordBuilder) GetCar() model.Car {
	b.car.Brand = "Ford"
	return b.car
}
