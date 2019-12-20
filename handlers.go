package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func addChannel(client *Client, data interface{}) {
	var channel Channel
	var message Message
	mapstructure.Decode(data, &channel)
	fmt.Printf("%#v\n", channel)
	// TODO: do RethinkDb logic here
	channel.Id = "73"
	message.Name = "add message"
	message.Data = channel
	client.send <- message
}
