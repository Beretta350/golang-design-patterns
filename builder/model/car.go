package model

import "fmt"

// Car represents the product being built.
type Car struct {
	Brand        string
	Model        string
	Engine       string
	Transmission string
	Seats        int
}

func (c *Car) ShowDetails() string {
	details := fmt.Sprintf("Brand: %s Model: %s Engine: %s Transmission: %s Seats: %d", c.Brand, c.Model, c.Engine, c.Transmission, c.Seats)
	return details
}
