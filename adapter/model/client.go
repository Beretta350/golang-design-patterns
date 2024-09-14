package model

import "fmt"

type Client struct {
}

func (c *Client) InsertUSBConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts USB connector into computer.")
	com.InsertIntoUSBPort()
}
