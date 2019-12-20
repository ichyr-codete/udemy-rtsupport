package main1

import (
	"encoding/json"
	"fmt"

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

func main() {
	raw := []byte(`{"name":"add channel", "data": {"name":"Hardware support"}}`)
	var data Message
	err := json.Unmarshal(raw, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", data)
	if data.Name == "add channel" {
		channel, err := addChannel(data.Data)
		if err != nil {
			fmt.Println(err)
			return
		}
		var sendMessage Message
		sendMessage.Name = "channel add"
		sendMessage.Data = channel
		sendRawMessage, err := json.Marshal(sendMessage)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(sendRawMessage))
	}
}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	// replaced by mapstructure that does type checking
	// channelMap := data.(map[string]interface{})
	// channel.Name = channelMap["name"].(string)
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}
	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	return channel, nil
}
