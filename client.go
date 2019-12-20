package main

import (
	"github.com/gorilla/websocket"
)

// FindHandler ...
type FindHandler func(string) (Handler, bool)

// Message ...
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// Client ...
type Client struct {
	send        chan Message
	socket      *websocket.Conn
	findHandler FindHandler
}

// Write ...
func (client *Client) Write() {
	for msg := range client.send {
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

// Read ...
func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

// NewClient ...
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		send:        make(chan Message),
		socket:      socket,
		findHandler: findHandler,
	}
}
