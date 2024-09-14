package factory

import (
	"fmt"

	"github.com/Beretta350/golang-design-patterns/factory/model"
)

func GetComputer(computerType string) (model.Computer, error) {
	var err error
	var comp model.Computer

	switch computerType {
	case "Gamer":
		comp = model.NewGamerComputer()
	case "Server":
		comp = model.NewServerComputer()
	default:
		err = fmt.Errorf("invalid computer type")
	}

	return comp, err
}
