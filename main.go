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

func main() {
	// fmt.Println("Hello, Ivan")
	http.HandleFunc("/", handler)
	fmt.Println("Server starting")
	http.ListenAndServe(":4000", nil)
	fmt.Println("Server started")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello from go!")
	fmt.Println(r.URL)
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msgType, msg, _ := socket.ReadMessage()
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}

	}
}
