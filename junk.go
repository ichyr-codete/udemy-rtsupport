package main

import (
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "127.0.0.1:28015",
		Database: "rtsupport",
	})
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	// Insert user
	user := User{
		Name: "Anonymous",
	}
	response, err := r.Table("user").
		Insert(user).
		RunWrite(session)

		// Update user
	user = User{
		Name: "John Moore",
	}
	response, _ = r.Table("user").
		Get("a8d9ecd2-dd10-4414-add8-47624a0eb481").Update(user).RunWrite(session)

	fmt.Println(response)
}
