package model

type GamerComputer struct {
	Computer
}

func NewGamerComputer() Computer {
	return &GamerComputer{
		Computer: &computer{
			OS:             "Windows",
			RAMSpace:       3200,
			HardDriveSpace: 40000,
		},
	}
}
