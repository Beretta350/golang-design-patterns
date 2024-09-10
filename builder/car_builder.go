package builder

import (
	"fmt"
)

// Car represents the product being built.
type Car struct {
	Brand        string
	Model        string
	Engine       string
	Transmission string
	Seats        int
}

func (c *Car) ShowDetails() string {
	details := fmt.Sprintf("Brand: %s\nModel: %s\nEngine: %s\nTransmission: %s\nSeats: %d\n", c.Brand, c.Model, c.Engine, c.Transmission, c.Seats)
	return details
}

// CarBuilder defines the steps for building a car.
type CarBuilder interface {
	SetModel(name string)
	SetEngine(engineType string)
	SetTransmission(transmissionType string)
	SetNumberOfSeats(numSeats int)
	GetCar() Car
}
