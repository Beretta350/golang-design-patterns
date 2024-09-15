package service

import (
	"context"
	"fmt"

	"github.com/Beretta350/golang-design-patterns/builder"
	"github.com/Beretta350/golang-design-patterns/builder/manufacturer"
	"github.com/Beretta350/golang-design-patterns/builder/model"
)

type CarService interface {
	CreateCar(ctx context.Context, brand string) model.Car
}

type carService struct {
	manager builder.CarManager
}

func NewCarService(m builder.CarManager) CarService {
	return &carService{manager: m}
}

func (serv *carService) CreateCar(ctx context.Context, brand string) model.Car {
	var builder builder.CarBuilder
	var car model.Car

	switch brand {
	case "Toyota":
		builder = &manufacturer.ToyotaBuilder{}
		car = serv.manager.
			SetBuilder(builder).
			SetModel("Corolla").
			SetEngine("1.8L").
			SetTransmission("Automatic").
			SetNumberOfSeats(5).BuildCar()

	case "Ford":
		builder = &manufacturer.FordBuilder{}
		serv.manager.SetBuilder(builder)
		car = serv.manager.
			SetBuilder(builder).
			SetModel("Mustang").
			SetEngine("5.0L V8").
			SetTransmission("Manual").
			SetNumberOfSeats(4).BuildCar()
	}

	fmt.Print(car.ShowDetails())

	return car
}
