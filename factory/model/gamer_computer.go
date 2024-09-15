package model

type GamerComputer struct {
	Computer
}

func NewGamerComputer() Computer {
	return &GamerComputer{
		Computer: NewComputer("Windows", 3200, 40000),
	}
}
