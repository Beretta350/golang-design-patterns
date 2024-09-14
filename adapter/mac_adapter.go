package adapter

import (
	"fmt"

	"github.com/Beretta350/golang-design-patterns/adapter/model"
)

type MacAdapter struct {
	MacMachine *model.Mac
}

func (m *MacAdapter) InsertIntoUSBPort() {
	fmt.Println("Mac adapter converts USB to USB-C.")
	m.MacMachine.InsertIntoUSBCPort()
}
