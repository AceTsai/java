package router

import (
	"github.com/gin-gonic/gin"
	"im/api"
	"im/middleware"
)

func Router() {
	engine := gin.Default()
	engine.Use(middleware.Cors())

	v1 := engine.Group("/api/v1")
	v1.POST("/auth/login", api.Login)
	v1.POST("/users/setting", api.Info)
	v1.POST("/talk/list", api.List)
	engine.Run(":9503")
}
