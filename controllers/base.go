package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 封装所有类的基类

type BaseController struct {
	*gin.Context
	gin.H
}

func (c *BaseController) response(msg string, data string) {
	// if len(data) < 1 {
	// 	data = "null"
	// }
	fmt.Println("data", data)
	c.JSON(http.StatusOK, gin.H{"success": true, "code": 200, "msg": msg, "data": data})

}
