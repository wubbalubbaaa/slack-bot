package slackbot

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/slack-go/slack"
)

type Reply struct {
}

var _ Handler = (*Reply)(nil)

const keywordsMap = `{
	"叶子":"简单整点那个",
	"猫猫":"烧杯猫猫之烧杯猫猫",
	"Bad bot":"没有教养的东西，怎么说话的",
	"bat bot":"嗯嗯",
	"宇宙":"你说宇宙很大很大，地球就是宇宙的一粒沙，我们人类连一粒沙都没有",
	"引流狗":"原来如此，原来如此，太谢谢了，我已经完全玩明白了",
	"神友就":"原来如此，原来如此，太谢谢了，我已经完全玩明白了",
	"t退役军人的苦逼":"t退役军人的苦逼，福建省人民医院骨科的庸医，陈鹏，石树培，林海仙，害的我出了医疗事故，把我搞得半死不活。在做腰椎手术之前左腿虽然很痛伸不直，右腿的大腿内侧 也痛但是能伸直。而现在是腿一边长一边短，右腰一直痛，关节和全身都痛而软，这样的庸医留着有何用",
	"神系":"唉，神友文化圈真的没落了",
	"臭逼":"好臭！嗯~~ 好臭！",
	"😅":"都让你流完了",
	"沙东人":"沙东人肿莫你了，触摸谁呢？我受你这气？",
	"山东人":"沙东人肿莫你了，触摸谁呢？我受你这气？",
	"日本人":"还好，一切如愿以偿",
	"day 0":"唉,共导吧",
	"白妈妈":"爱不爱妈妈",
	"可以吗":"不阔以!",
	"不会吧不会吧":"我不会你妈隔壁",
	"红迪":"哈哈红迪狗都不看,我上红迪只封人"
	}`

func (reply *Reply) OnMessge(client *slack.Client, msg *slack.MessageEvent) error {
	if msg.BotID != "" {
		return BotMessageErr
	}
	msgstr := strings.Trim(msg.Text, " ")
	msgstr = strings.Trim(msgstr, "\n")
	keywords, err := JsonToMap(keywordsMap)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range keywords {
		if idx := strings.Index(msg.Text, k); idx != -1 {
			client.PostMessage(msg.Channel, slack.MsgOptionAsUser(true), slack.MsgOptionText(v, false))
		}

	}
	return nil
}

func (reply *Reply) OnError(client *slack.Client, msg *slack.RTMError) error {
	log.Println("error", msg)
	return nil
}

func (reply *Reply) OnConnect(client *slack.Client, msg *slack.ConnectedEvent) error {
	log.Println("on connenct")
	return nil
}

func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		log.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	return m, nil
}
