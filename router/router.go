package router

import (
	"github.com/gin-gonic/gin"

	"github.com/luxingwen/Bass/auth"
	"github.com/luxingwen/Bass/controllers"
)

func Router() *gin.Engine {
	r := gin.New()

	// 路由
	v1 := r.Group("/api")
	// controllers.UserRegister(v1.Group("/classes"))
	controllers.UserRegister(v1)

	r.Use(gin.Logger())
	r.Use(auth.AuthMiddleware(false))
	return r
}
