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
		fmt.Println(err)
		return
	}
}
