package manager

import (
	"github.com/Beretta350/golang-design-patterns/builder"
	"github.com/Beretta350/golang-design-patterns/builder/model"
)

// CarManager is responsible for managing the car building process.
type CarManager struct {
	builder builder.CarBuilder
}

func (d *CarManager) SetBuilder(b builder.CarBuilder) {
	d.builder = b
}

func (d *CarManager) BuildCar(model, engine, transmission string, seats int) model.Car {
	d.builder.SetModel(model)
	d.builder.SetEngine(engine)
	d.builder.SetTransmission(transmission)
	d.builder.SetNumberOfSeats(seats)
	return d.builder.GetCar()
}
