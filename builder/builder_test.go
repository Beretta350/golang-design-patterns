package builder_test

import (
	"context"
	"testing"

	"github.com/Beretta350/golang-design-patterns/builder"
	"github.com/Beretta350/golang-design-patterns/builder/model"
	"github.com/Beretta350/golang-design-patterns/builder/service"
	"github.com/stretchr/testify/assert"
)

func TestCarBuilderCreation(t *testing.T) {
	ctx := context.Background()

	man := builder.CarManager{}
	serv := service.NewCarService(man)

	toyotaCar := serv.CreateCar(ctx, "Toyota")
	fordCar := serv.CreateCar(ctx, "Ford")

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
