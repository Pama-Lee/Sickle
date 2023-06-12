package server

import "github.com/gin-gonic/gin"

// serverRoute 注册路由
func serverRoute(c *gin.RouterGroup) {
	// 注册路由, GET和POST
	c.GET("/ping", ping)
}

// ping 检查服务是否正常
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "pong",
	})

}
