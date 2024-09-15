package observer

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}
