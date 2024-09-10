package service

import (
	"context"
	"fmt"

	"github.com/Beretta350/golang-design-patterns/builder"
	"github.com/Beretta350/golang-design-patterns/builder/manager"
	"github.com/Beretta350/golang-design-patterns/builder/manufacturer"
	"github.com/Beretta350/golang-design-patterns/builder/model"
)

type CarService interface {
	CreateCar(ctx context.Context, brand string) model.Car
}

type carService struct {
	manager manager.CarManager
}

func NewCarService(m manager.CarManager) CarService {
	return &carService{manager: m}
}

func (serv *carService) CreateCar(ctx context.Context, brand string) model.Car {
	var builder builder.CarBuilder
	var car model.Car

	switch brand {
	case "Toyota":
		builder = &manufacturer.ToyotaBuilder{}
		serv.manager.SetBuilder(builder)
		car = serv.manager.BuildCar("Corolla", "1.8L", "Automatic", 5)
	case "Ford":
		builder = &manufacturer.FordBuilder{}
		serv.manager.SetBuilder(builder)
		car = serv.manager.BuildCar("Mustang", "5.0L V8", "Manual", 4)
	}

	fmt.Print(car.ShowDetails())

	return car
}
