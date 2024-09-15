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

func NewComputer(os string, ram, hd int) Computer {
	return &computer{OS: os, RAMSpace: ram, HardDriveSpace: hd}
}

func (c *computer) SetOS(name string) {
	c.OS = name
}

func (c *computer) SetRAMSpace(ram int) {
	c.RAMSpace = ram
}

func (c *computer) SetHDSpace(hd int) {
	c.HardDriveSpace = hd
}

func (c *computer) GetDetails() string {
	return fmt.Sprintf("OS: %s, RAM: %dGB, Hard Drive: %dGB", c.OS, c.RAMSpace, c.HardDriveSpace)
}

func (c *computer) GetOS() string {
	return c.OS
}

func (c *computer) GetRAMSpace() int {
	return c.RAMSpace
}

func (c *computer) GetHDSpace() int {
	return c.HardDriveSpace
}
