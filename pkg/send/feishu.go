package send

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/myoperator/messageoperator/pkg/sysconfig"
	"github.com/myoperator/messageoperator/pkg/workqueue"
	"net/http"
)

// FeishuClient 邮件发送器
type FeishuClient struct {
	client *http.Client
}

// Send 发送邮件
func (fc *FeishuClient) Send(obj *workqueue.QueueResource) error {
	var template map[string]interface{}
	if SysConfig1.Feishu.Type == "card" {
		template = setCardMessageTemplate(obj.Kind, obj.Name, obj.Namespace, string(obj.EventType))
	}

	// 将卡片消息内容转换为 JSON
	payload, err := json.Marshal(template)
	if err != nil {
		fmt.Println("Error marshaling card message:", err)
		return err
	}

	// 发送 POST 请求到飞书 Webhook URL
	resp, err := http.Post(fmt.Sprintf("https://open.feishu.cn/open-apis/bot/v2/hook/%s", SysConfig1.Feishu.Webhook), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response received. Status:", resp.Status)
		return err
	}
	return nil
}

// NewSender 创建邮件发送器
func NewFeishuClient() *FeishuClient {
	return &FeishuClient{
		client: http.DefaultClient,
	}
}

func setCardMessageTemplate(kind, name, namespace, eventType string) map[string]interface{} {

	// 卡片模版内容
	template := map[string]interface{}{
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
					"content": fmt.Sprintf("集群内有资源发生变动事件，资源[%s]，[%s/%s]，事件[%s], 如有疑问请尽快处理。\n<at id=all></at>", kind, name, namespace, eventType),
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

	return template
}

var (
	// 卡片模版内容
	cardMessageTemplate = map[string]interface{}{
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
)
