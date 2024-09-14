package model

type ServerComputer struct {
	Computer
}

func NewServerComputer() Computer {
	return &ServerComputer{
		Computer: &computer{
			OS:             "Linux",
			RAMSpace:       128000,
			HardDriveSpace: 4000000,
		},
	}
}
