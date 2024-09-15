package model

type ServerComputer struct {
	Computer
}

func NewServerComputer() Computer {
	return &ServerComputer{
		Computer: NewComputer("Linux", 128000, 4000000),
	}
}
