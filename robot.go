package slackbot

import (
	"github.com/slack-go/slack"
)

const BotID = "U03HLM3GJDN"

type Handler interface {
	OnConnect(client *slack.Client, msg *slack.ConnectedEvent) error
	OnMessge(client *slack.Client, msg *slack.MessageEvent) error
	OnError(client *slack.Client, msg *slack.RTMError) error
}

type Robot struct {
	client   *slack.Client
	rtm      *slack.RTM
	handlers []Handler
}

func NewRobot(token string, handlers []Handler) *Robot {
	client := slack.New(token)
	rtm := client.NewRTM()
	return &Robot{client: client, rtm: rtm, handlers: handlers}
}

func (rb *Robot) Run() {
	go rb.rtm.ManageConnection()
	for {
		select {
		case msg := <-rb.rtm.IncomingEvents:
			// log.Println("Event Received:", reflect.TypeOf(msg.Data))
			for i := range rb.handlers {
				if v, ok := msg.Data.(*slack.ConnectedEvent); ok {
					rb.handlers[i].OnConnect(rb.client, v)
				}
				if v, ok := msg.Data.(*slack.MessageEvent); ok {
					rb.handlers[i].OnMessge(rb.client, v)
				}
				if v, ok := msg.Data.(*slack.RTMError); ok {
					rb.handlers[i].OnError(rb.client, v)
				}
				// default
				// fmt.Println(msg.Data)
			}
		}
	}
}
