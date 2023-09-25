package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/919c368d-bbbf-4858-b788-0b78a4389d45" // 替换为您的飞书机器人 Webhook URL

	// 构建卡片消息内容
	cardMessage := map[string]interface{}{
		"msg_type": "interactive",
		"card": map[string]interface{}{
			"config": map[string]interface{}{
				"wide_screen_mode": true,
			},
			"header": map[string]interface{}{
				"template": "blue",
				"title": map[string]interface{}{
					"content": "集群内部资源变动通知",
					"tag":     "plain_text",
				},
			},
			"elements": []interface{}{
				map[string]interface{}{
					"tag":     "markdown",
					"content": fmt.Sprintf("集群内有资源发生变动事件，资源[%s]，事件[%s]，如有疑问请尽快处理。\n<at id=all></at>", "deployment", "add"),
				},
				map[string]interface{}{
					"tag": "action",
					"actions": []interface{}{
						map[string]interface{}{
							"tag": "button",
							"text": map[string]interface{}{
								"content": "按钮",
								"tag":     "plain_text",
							},
							"type": "primary",
							"url":  "https://www.feishu.cn",
						},
					},
				},
			},
		},
	}

	// 将卡片消息内容转换为 JSON
	payload, err := json.Marshal(cardMessage)
	if err != nil {
		fmt.Println("Error marshaling card message:", err)
		return
	}

	// 发送 POST 请求到飞书 Webhook URL
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response received. Status:", resp.Status)
		return
	}

	fmt.Println("Card message sent successfully!")
}
