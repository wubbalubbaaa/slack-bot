package main

import (
	"fmt"
	"os"
	slackbot "slack-bot"
)

var token = os.Getenv("SLACK_TOKEN")

func main() {
	fmt.Println(token)
	ch := make(chan int)
	handlers := []slackbot.Handler{&slackbot.Poll{}, &slackbot.Reply{}}
	rb := slackbot.NewRobot(token, handlers)
	go rb.Run()
	ch <- 1
}
