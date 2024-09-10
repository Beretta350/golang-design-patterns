package builder_test

import (
	"testing"

	"github.com/Beretta350/golang-design-patterns/builder"
	"github.com/stretchr/testify/assert"
)

func TestCarBuilderCreation(t *testing.T) {
	toyotaBuilder := &builder.ToyotaBuilder{}
	fordBuilder := &builder.FordBuilder{}

	manager := builder.CarManager{}

	manager.SetBuilder(toyotaBuilder)
	toyotaCar := manager.BuildCar("Corolla", "1.8L", "Automatic", 5)
	t.Log(toyotaCar.ShowDetails())

	manager.SetBuilder(fordBuilder)
	fordCar := manager.BuildCar("Mustang", "5.0L V8", "Manual", 4)
	t.Log(fordCar.ShowDetails())

	toyotaExpectedCar := builder.Car{
		Brand:        "Toyota",
		Model:        "Corolla",
		Engine:       "1.8L",
		Transmission: "Automatic",
		Seats:        5,
	}

	fordExpectedCar := builder.Car{
		Brand:        "Ford",
		Model:        "Mustang",
		Engine:       "5.0L V8",
		Transmission: "Manual",
		Seats:        4,
	}

	assert.Equal(t, toyotaCar, toyotaExpectedCar)
	assert.Equal(t, fordCar, fordExpectedCar)
}
