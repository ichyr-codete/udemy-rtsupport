package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting")
	http.ListenAndServe(":4000", nil)
	fmt.Println("Server started")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var inMessage Message
		var outMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}
		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage = Message{"error", err}
				if err := socket.WriteJSON(outMessage); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			{
				go subscribeChannel(socket)
			}
		}
	}
}

func addChannel(data interface{}) error {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	return nil
}

func subscribeChannel(socket *websocket.Conn) {
	// TODO: rethinkDB Query / changefeed
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add", Channel{"1", "Software Support"}}
		socket.WriteJSON(message)
	}

}
