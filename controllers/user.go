package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/luxingwen/Bass/controllers"
)

// 用户操作接口
// 1.登陆 2.注册 3.找回密码 4.第三方登陆 5.重置密码
func UserRegister(router *gin.RouterGroup) {
	v1.POST("/login", controllers.Login)
	v1.POST("/Create", controllers.Create)

}

// 测试
// func Login(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"success": true, "code": 200, "msg": "hello"})
// }

// 登陆
func Login(c *gin.Context) {
	bodyRaw, _ := c.GetRawData()

	m := make(map[string]interface{})
	err := json.Unmarshal(bodyRaw, &m)
	if err != nil {
		log.Errorf("GetRawData err:%s ", err)
		c.JSON(http.StatusOK, gin.H{"success": false, "code": 200, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "code": 200, "msg": "hello"})
}

// 创建用户
func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "code": 200, "msg": "hello"})
}
