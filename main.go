package main

import (
	"log"
	"net/http"
	"time"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

// Channel ...
type Channel struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
}

// User ...
type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

// ChannelMessage ...
type ChannelMessage struct {
	Id        string    `gorethink:"id,omitempty"`
	ChannelId string    `gorethink:"channelId"`
	Body      string    `gorethink:"body"`
	Author    string    `gorethink:"author"`
	CreatedAt time.Time `gorethink:"createdAt"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "127.0.0.1:28015",
		Database: "rtsupport",
	})
	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)

	// chanel handlers
	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)

	// user handlers
	router.Handle("user edit", editUser)
	router.Handle("user subscribe", subscribeUser)
	router.Handle("user unsubscribe", unsubscribeUser)

	// message handlers
	router.Handle("message add", addChannelMessage)
	router.Handle("message subscribe", subscribeChannelMessage)
	router.Handle("message unsubscribe", unsubscribeChannelMessage)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
