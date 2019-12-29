package main

import (
	"github.com/gorilla/websocket"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
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
	send         chan Message
	socket       *websocket.Conn
	findHandler  FindHandler
	session      *r.Session
	stopChannels map[int]chan bool
}

// NewStopChannel ...
func (client *Client) NewStopChannel(stopKey int) chan bool {
	client.StopForKey(stopKey)
	stop := make(chan bool)
	client.stopChannels[stopKey] = stop
	return stop
}

// StopForKey ...
func (client *Client) StopForKey(stopKey int) {
	if ch, found := client.stopChannels[stopKey]; found {
		ch <- true
		delete(client.stopChannels, stopKey)
	}
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

// Close ...
func (client *Client) Close() {
	for _, stopChannel := range client.stopChannels {
		stopChannel <- true
	}
	close(client.send)
}

// NewClient ...
func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) *Client {
	return &Client{
		send:         make(chan Message),
		socket:       socket,
		findHandler:  findHandler,
		session:      session,
		stopChannels: make(map[int]chan bool),
	}
}
