package model

import (
	"fmt"

	"github.com/Beretta350/golang-design-patterns/observer"
)

type Item struct {
	observerList []observer.Observer
	name         string
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) Register(o observer.Observer) {
	fmt.Printf("Adding observer %s in the %s observer's list\n", o.GetID(), i.name)
	i.observerList = append(i.observerList, o)
}

func (i *Item) DeRegister(o observer.Observer) {
	fmt.Printf("Removing observer %s from %s observer's list\n", o.GetID(), i.name)
	i.observerList = removeFromList(i.observerList, o)
}

func (i *Item) NotifyAll() {
	fmt.Printf("Notifying all observers about %s\n", i.name)
	for _, observer := range i.observerList {
		observer.Update(i.name)
	}
}

func removeFromList(observerList []observer.Observer, removing observer.Observer) []observer.Observer {
	for i, observer := range observerList {
		if removing.GetID() == observer.GetID() {
			return append(observerList[:i], observerList[i+1:]...)
		}
	}
	return observerList
}
