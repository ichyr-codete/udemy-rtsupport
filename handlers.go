package main

import (
	"github.com/mitchellh/mapstructure"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	go func() {
		err = r.Table("channel").Insert(channel).Exec(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
		}
	}()
}

func subscribeChannel(client *Client, data interface{}) {
	// we know subscription to RethinkDB change feed  is blocking operation
	// so we spawn separate goroutine for it
	go func() {
		cursor, err := r.Table("channels").
			Changes(r.ChangesOpts{IncludeInitial: true}).
			Run(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}
		var change r.ChangeResponse
		for cursor.Next(&change) {
			if change.NewValue != nil && change.OldValue == nil {
				client.send <- Message{"channel add", change.NewValue}
			}
		}
	}()
}
