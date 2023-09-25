package httpserver

import (
	"github.com/gin-gonic/gin"
	. "github.com/myoperator/messageoperator/pkg/send"
	"k8s.io/klog/v2"
)

type SendRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// SendTo 发送信息逻辑
func SendTo(c *gin.Context) {
	var r *SendRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		klog.Error("bind json err!")
		c.JSON(400, gin.H{"error": err})
		return
	}
	s := NewEmailSender()
	err := s.Send(r.Title, r.Content)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"ok": "ok"})
	return
}
