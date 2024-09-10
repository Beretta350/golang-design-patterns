package manufacturer

import "github.com/Beretta350/golang-design-patterns/builder/model"

// ToyotaBuilder is a concrete builder for Toyota cars.
type ToyotaBuilder struct {
	car model.Car
}

func (b *ToyotaBuilder) SetModel(name string) {
	b.car.Model = name
}

func (b *ToyotaBuilder) SetEngine(engineType string) {
	b.car.Engine = engineType
}

func (b *ToyotaBuilder) SetTransmission(transmissionType string) {
	b.car.Transmission = transmissionType
}

func (b *ToyotaBuilder) SetNumberOfSeats(numSeats int) {
	b.car.Seats = numSeats
}

func (b *ToyotaBuilder) GetCar() model.Car {
	b.car.Brand = "Toyota"
	return b.car
}
