package server

import (
	"Hooker/entity"
	"Hooker/log"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Start 启动http服务
func Start(config *entity.Config) {
	// 启动http服务
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	gin.DefaultWriter = log.Writer

	// 注册路由
	registerRouter(r)

	portStr := strconv.Itoa(config.Server.Port)

	log.Info("Server is running on port " + portStr)
	// 启动服务
	err := r.Run(":" + portStr)
	if err != nil {
		log.Error(err)
		panic(err)
	}

}
