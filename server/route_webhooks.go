package server

import (
	"Sickle/entity/databaseEntity"
	"Sickle/entity/request"
	"Sickle/log"
	"Sickle/tool"

	"github.com/gin-gonic/gin"
)

// WebhooksRoute 注册路由
func WebhooksRoute(c *gin.RouterGroup) {
	// 注册路由, GET和POST
	c.POST("/post/:token", post)
	c.GET("/get", get)

}

// post 接收webhooks请求
func post(c *gin.Context) {

	// 获取token
	token := c.Param("token")
	if token == "" {
		tool.BadRequestResponse(c, "token is empty")
		return
	}

	// 验证token
	if !VerifyUUID(token) {
		tool.BadRequestResponse(c, "token is invalid")
		return
	}

	// 获取请求body
	body, err := c.GetRawData()
	if err != nil {
		tool.BadRequestResponse(c, "Failed to get request body")
		log.Error("Failed to get request body: ", err)
		return
	}

	var req request.WebhookRequest
	req.Body = string(body)

	var webhookLog databaseEntity.WebhooksLog
	webhookLog.Body = req.Body

}

// get 接收webhooks请求
func get(c *gin.Context) {

}
