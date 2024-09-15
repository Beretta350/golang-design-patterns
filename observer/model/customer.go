package model

import "fmt"

// The customer is the observer
type Customer struct {
	Name string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending a message to customer %s about %s\n", c.Name, itemName)
}

func (c *Customer) GetID() string {
	return c.Name
}
