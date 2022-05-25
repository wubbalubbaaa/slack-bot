package slackbot

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/slack-go/slack"
)

type Poll struct {
}

var _ Handler = (*Poll)(nil)

func (poll *Poll) OnMessge(client *slack.Client, msg *slack.MessageEvent) error {
	if msg.BotID != "" {
		return BotMessageErr
	}
	msgstr := strings.Trim(msg.Text, " ")
	msgstr = strings.Trim(msgstr, "\n")
	msgstr = strings.Join(strings.FieldsFunc(msgstr, unicode.IsSpace), "")
	if hasPrefix(msgstr, fmt.Sprintf("<@%s>投票：", BotID)) || hasPrefix(msgstr, fmt.Sprintf("<@%s>投票:", BotID)) {
		msgRef := slack.NewRefToMessage(msg.Channel, msg.Timestamp)
		client.AddReaction("white_check_mark", msgRef)
		client.AddReaction("x", msgRef)
	}
	return nil
}

func (poll *Poll) OnError(client *slack.Client, msg *slack.RTMError) error {
	log.Println("error", msg)
	return nil
}

func (poll *Poll) OnConnect(client *slack.Client, msg *slack.ConnectedEvent) error {
	log.Println("on connenct")
	return nil
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}
