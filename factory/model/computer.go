package model

import "fmt"

type Computer interface {
	SetOS(name string)
	SetRAMSpace(ram int)
	SetHDSpace(ram int)
	GetDetails() string
	GetOS() string
	GetRAMSpace() int
	GetHDSpace() int
}

type computer struct {
	OS             string
	RAMSpace       int
	HardDriveSpace int
}

func NewComputer(name, os string, ram, hd int) Computer {
	return &computer{OS: os, RAMSpace: ram, HardDriveSpace: hd}
}

// SetOS sets the OS of the computer
func (c *computer) SetOS(name string) {
	c.OS = name
}

// SetRAMSpace sets the RAM space of the computer
func (c *computer) SetRAMSpace(ram int) {
	c.RAMSpace = ram
}

// SetHDSpace sets the hard drive space of the computer
func (c *computer) SetHDSpace(hd int) {
	c.HardDriveSpace = hd
}

// GetDetails returns the details of the computer as a string
func (c *computer) GetDetails() string {
	return fmt.Sprintf("OS: %s, RAM: %dGB, Hard Drive: %dGB", c.OS, c.RAMSpace, c.HardDriveSpace)
}

// SetOS sets the OS of the computer
func (c *computer) GetOS() string {
	return c.OS
}

// GetRAMSpace Gets the RAM space of the computer
func (c *computer) GetRAMSpace() int {
	return c.RAMSpace
}

// GetHDSpace Gets the hard drive space of the computer
func (c *computer) GetHDSpace() int {
	return c.HardDriveSpace
}
