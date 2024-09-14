package factory_test

import (
	"fmt"
	"testing"

	"github.com/Beretta350/golang-design-patterns/factory"
	"github.com/stretchr/testify/assert"
)

func TestComputerFactory(t *testing.T) {
	gamer, _ := factory.GetComputer("Gamer")
	server, _ := factory.GetComputer("Server")

	fmt.Println(gamer.GetDetails())
	fmt.Println(server.GetDetails())

	assert.Equal(t, "Windows", gamer.GetOS())
	assert.Equal(t, 3200, gamer.GetRAMSpace())
	assert.Equal(t, 40000, gamer.GetHDSpace())

	assert.Equal(t, "Linux", server.GetOS())
	assert.Equal(t, 128000, server.GetRAMSpace())
	assert.Equal(t, 4000000, server.GetHDSpace())
}
