package adapter_test

import (
	"testing"

	"github.com/Beretta350/golang-design-patterns/adapter"
	"github.com/Beretta350/golang-design-patterns/adapter/model"
)

func TestMacAdapter(t *testing.T) {
	client := &model.Client{}
	dell := &model.Dell{}

	client.InsertUSBConnectorIntoComputer(dell)

	macMachine := &model.Mac{}
	macMachineAdapter := &adapter.MacAdapter{
		MacMachine: macMachine,
	}

	client.InsertUSBConnectorIntoComputer(macMachineAdapter)
}
