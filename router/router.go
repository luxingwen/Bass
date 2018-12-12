package router

import (
	"github.com/gin-gonic/gin"

	"github.com/luxingwen/Bass/auth"
	"github.com/luxingwen/Bass/controllers"
)

func Router() *gin.Engine {
	r := gin.New()

	// 路由
	v1 := r.Group("/api/v1")

	// 数据库操作控制器
	v1.GET("/hello", controllers.Hellos)
	v1.GET("/classes/:tableName", controllers.GetAllClasses)
	v1.POST("/classes/:tableName", controllers.PostClasses)
	v1.GET("/classes/:tableName/:id", controllers.GetClasses)
	v1.PUT("/classes/:tableName/:id", controllers.PutClasses)
	v1.DELETE("/classes/:tableName/:id", controllers.DeleteClasses)
	// controllers.UserRegister(v1)

	// 用户操作接口
	// 1.登陆 2.注册 3.找回密码 4.第三方登陆 5.重置密码
	// v1.POST("/login", controllers.Login)
	// v1.POST("/Create", controllers.Create)

	r.Use(gin.Logger())
	r.Use(auth.AuthMiddleware(false))
	return r
}
