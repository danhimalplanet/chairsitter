package main

import (
"github.com/BurntSushi/toml"
	"./event"
	"./slackuserid"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
		"strings"
	"time"
)

type Event event.Event

var Events []Event

var balance int

var conf Config

type Config struct {
        General  general
}

type general struct {
        Slacktoken string
}


func main() {

        if _, err := toml.DecodeFile("./secrets.toml", &conf); err != nil {
                fmt.Println(err)
        }

	slacktoken := conf.General.Slacktoken

	endpoints := map[string]string{
		"peets":     "starbucks",
		"starbucks": "peets",
	}

	url := fmt.Sprintf("http://localhost:4242/api/messages")

	balance = 1

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
			return
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Got a fatal error: ", err)
			return
		}

		incomingMessages := func(req *http.Request, resp *http.Response) []Event {
			defer resp.Body.Close()
			if err := json.NewDecoder(resp.Body).Decode(&Events); err != nil {
				log.Println(err)
			}
			return Events
		}(req, resp)

		for _, incomingMessage := range incomingMessages {
			t1 := incomingMessage.Text

			// Incog land
			if incomingMessage.Gateway == "peets" {
				if balance > 0 {
					balance = balance - 1
					event.SendMessage(incomingMessage.Username, t1, endpoints[incomingMessage.Gateway])
				} else {
					event.SendMessage("alert", "You need 1 token to chat.", incomingMessage.Gateway)
				}
			}

			// Slack side of things
			if incomingMessage.Gateway == "starbucks" {
				fmt.Printf("From Slack side balance is %d\n", balance)
				chatter := slackuserid.SlackIDtoUsername(slacktoken, incomingMessage.Userid)
				if strings.Contains(incomingMessage.Text, "incog") {
					balance = 2
				}
				event.SendMessage(chatter, incomingMessage.Text, endpoints[incomingMessage.Gateway])
			}
		}

		time.Sleep(2 * time.Second)

	}

}
