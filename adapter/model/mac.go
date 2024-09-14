package model

import "fmt"

type Mac struct{}

func (m *Mac) InsertIntoUSBCPort() {
	fmt.Println("USB-C connected into a MAC.")
}
