package builder_test

import (
	"fmt"
	"testing"

	"github.com/Beretta350/golang-design-patterns/builder"
	"github.com/Beretta350/golang-design-patterns/builder/manufacturer"
	"github.com/Beretta350/golang-design-patterns/builder/model"
	"github.com/stretchr/testify/assert"
)

func TestCarBuilderCreation(t *testing.T) {
	manager := builder.CarManager{}
	var builder builder.CarBuilder

	builder = &manufacturer.ToyotaBuilder{}

	toyotaCar := manager.SetBuilder(builder).
		SetModel("Corolla").
		SetEngine("1.8L").
		SetTransmission("Automatic").
		SetNumberOfSeats(5).BuildCar()

	builder = &manufacturer.FordBuilder{}

	fordCar := manager.SetBuilder(builder).
		SetModel("Mustang").
		SetEngine("5.0L V8").
		SetTransmission("Manual").
		SetNumberOfSeats(4).BuildCar()

	fmt.Println(toyotaCar.ShowDetails())
	fmt.Println(fordCar.ShowDetails())

	toyotaExpectedCar := model.Car{
		Brand:        "Toyota",
		Model:        "Corolla",
		Engine:       "1.8L",
		Transmission: "Automatic",
		Seats:        5,
	}

	fordExpectedCar := model.Car{
		Brand:        "Ford",
		Model:        "Mustang",
		Engine:       "5.0L V8",
		Transmission: "Manual",
		Seats:        4,
	}

	assert.Equal(t, toyotaCar, toyotaExpectedCar)
	assert.Equal(t, fordCar, fordExpectedCar)
}
