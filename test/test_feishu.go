package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

type SendMessage struct {
	MsgType string            `json:"msg_type"`
	Content map[string]string `json:"content"`
}

func sendMsg(apiUrl, msg string) {
	// json
	contentType := "application/json"
	// data
	//sendData := `{
	//	"msg_type": "text",
	//	"content": {"text": "` + "消息通知:" + msg + `"}
	//}`

	sendData := `{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "template": "blue",
    "title": {
      "tag": "plain_text",
      "content": "集群内部资源变动通知"
    }
  },
  "elements": [
    {
      "tag": "markdown",
      "content": "集群内有资源发生变动事件，资源[]，事件[]，如有疑问请尽快处理。\n<at id=all></at>"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "这是跳转按钮"
          },
          "type": "primary",
          "url": "https://feishu.cn"
        }
      ]
    }
  ]
}`

	//aa := &SendMessage{
	//	MsgType: "text",
	//	Content: map[string]string{},
	//}
	//aa.Content["text"] = "dddddsssss"
	//
	//aaa, err := json.Marshal(aa)
	//if err != nil {
	//	log.Fatal(err)
	//}
	// request
	result, err := http.Post(apiUrl, contentType, strings.NewReader(sendData))
	//result, err := http.Post(apiUrl, contentType, bytes.NewReader(aaa))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()

}

func main() {
	// webhook地址
	var webhookUrl string
	// 消息内容
	var message string

	flag.StringVar(&webhookUrl, "u", "https://open.feishu.cn/open-apis/bot/v2/hook/xxxx", "飞书webhook地址")
	flag.StringVar(&message, "s", "ddddddd", "需要发送的消息内容")

	flag.Parse()
	flag.Usage()
	sendMsg(webhookUrl, message)
}
