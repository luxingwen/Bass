package router

import (
	"github.com/gin-gonic/gin"

	"github.com/luxingwen/Bass/auth"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(auth.AuthMiddleware(false))
	return r
}
