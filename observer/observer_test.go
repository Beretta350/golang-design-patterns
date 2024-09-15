package observer_test

import (
	"testing"

	"github.com/Beretta350/golang-design-patterns/observer/model"
)

func TestObserver(t *testing.T) {
	subject := model.NewItem("Golang language")

	observer1 := &model.Customer{Name: "Robert Griesemer"}
	observer2 := &model.Customer{Name: "Rob Pike"}
	observer3 := &model.Customer{Name: "Ken Thompson"}

	subject.Register(observer1)
	subject.Register(observer2)
	subject.Register(observer3)

	subject.NotifyAll()

	subject.DeRegister(observer3)
	subject.DeRegister(observer2)

	subject.NotifyAll()

	subject.Register(observer2)
	subject.Register(observer3)
}
