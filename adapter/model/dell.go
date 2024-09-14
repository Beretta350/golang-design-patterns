package model

import "fmt"

type Dell struct{}

func (w *Dell) InsertIntoUSBPort() {
	fmt.Println("USB connected into a Dell.")
}
