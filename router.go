package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Handler ...
type Handler func(*Client, interface{})

// Router ...
type Router struct {
	rules map[string]Handler
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

// Handle ...
func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

// ServeHTTP ...
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	client := NewClient(socket)

	// methods below need to run in separate goroutines
	// lets spawn Write method in it's own goroutine and
	// let Read method use goroutine that was created when
	// ServeHTTP method was called
	go client.Write()
	client.Read()
}
