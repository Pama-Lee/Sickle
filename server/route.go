package server

import "github.com/gin-gonic/gin"

// registerRouter 注册路由
func registerRouter(r *gin.Engine) {
	// 注册路由, GET和POST
	server := r.Group("/server")
	{
		serverRoute(server)
	}

}
