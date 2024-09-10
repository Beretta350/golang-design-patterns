package builder

// CarManager is responsible for managing the car building process.
type CarManager struct {
	builder CarBuilder
}

func (d *CarManager) SetBuilder(b CarBuilder) {
	d.builder = b
}

func (d *CarManager) BuildCar(model, engine, transmission string, seats int) Car {
	d.builder.SetModel(model)
	d.builder.SetEngine(engine)
	d.builder.SetTransmission(transmission)
	d.builder.SetNumberOfSeats(seats)
	return d.builder.GetCar()
}
