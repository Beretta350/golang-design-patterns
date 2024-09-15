package observer

type Observer interface {
	Update(string)
	GetID() string
}
