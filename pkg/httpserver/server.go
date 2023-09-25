package httpserver

import "github.com/gin-gonic/gin"

func HttpServer() error {
	router := gin.New()

	register(router)
	err := router.Run(":8888")
	// TODO: 端口需要可配置
	if err != nil {
		return err
	}
	return nil
}

func register(router *gin.Engine) {

	r := router.Group("/v1")
	{
		// 对外发送email通知
		r.POST("/send", SendTo)
	}
}
