package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"
	"strings"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/luxingwen/Bass/models"
)

// 数据库操作控制器
// func UserRegister(router *gin.RouterGroup) {
// 	router.GET("/hello", Hellos)
// 	router.GET("/classes/:tableName", GetAllClasses)
// 	router.POST("/classes/:tableName", PostClasses)
// 	router.GET("/classes/:tableName/:id", GetClasses)
// 	router.PUT("/classes/:tableName/:id", PutClasses)
// 	router.DELETE("/classes/:tableName/:id", DeleteClasses)
// }

// 测试
func Hellos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "code": 200, "msg": "hello"})

	// var k BaseController
	// k.response("错误", "test")

}

// 删除一条记录
func DeleteClasses(c *gin.Context) {
	db2 := models.GetDB()
	tableName := c.Param("tableName")
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return

	}

	type tableId struct {
		Id int
	}
	k := &tableId{Id: id}
	affected, err := db2.Table(tableName).Delete(k)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "code": 101, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "code": 200, "msg": "ok", "status": affected})
}

// 修改一条记录
func PutClasses(c *gin.Context) {
	db := models.GetDB()

	tableName := c.Param("tableName")
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	content, _ := c.GetRawData()

	m := make(map[string]interface{})
	err = json.Unmarshal(content, &m)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101, "msg": err})
		return
	}

	affected, err := db.Table(tableName).Where("id=?", id).Update(m)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "code": 101, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "code": 200, "msg": "ok", "data": affected})
}

// 获取一行记录
func GetClasses(c *gin.Context) {
	db := models.GetDB()
	tableName := c.Param("tableName")
	id := c.Param("id")

	limit := 1
	offset := 0

	type result map[string]*string
	var results []result

	db.Table(tableName).AllCols().Where("id=?", id).Limit(limit, offset).Find(&results)

	if len(results) < 1 {
		c.JSON(http.StatusOK, gin.H{"success": false, "code": 101, "msg": "未找到记录"})
		log.Errorf("%s未找到记录", tableName)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

// 增加记录
/**
问题修复
1.字段类型匹配可能存在问题
**/
func PostClasses(c *gin.Context) {
	db2 := models.GetDB()
	tableName := c.Param("tableName")
	bodyRaw, _ := c.GetRawData()

	m := make(map[string]interface{})
	err := json.Unmarshal(bodyRaw, &m)

	msg := "ok"

	// /*  INSERT INTO `abc_models` (`id`, `created_at`, `updated_at`, `deleted_at`, `slug`, `title`, `description`, `body`, `author_id`) VALUES (NULL, NULL, NULL, NULL, NULL, '334', NULL, NULL, NULL);
	sqlf := "INSERT INTO `%s`(%s) VALUES (%s)"

	var field, value string
	var fields, values []interface{}

	// 处理列与值
	for k, v := range m {
		// 数组拿到所有列名称
		fields = append(fields, fmt.Sprintf("%s", k))
		// fmt.Println(v, reflect.TypeOf(v)) //了解传入的数据类型
		// 强制类型转换
		switch v := v.(type) {
		case string:
			values = append(values, fmt.Sprintf("'%s'", v))
		case int64:
			values = append(values, fmt.Sprintf("%d", v))
		case float64:
			values = append(values, float64(v))
		}
	}
	// 数组逗号分隔为字符串
	field = strings.Replace(strings.Trim(fmt.Sprint(fields), "[]"), " ", ",", -1)
	value = strings.Replace(strings.Trim(fmt.Sprint(values), "[]"), " ", ",", -1)

	sql := fmt.Sprintf(sqlf, tableName, field, value)
	fmt.Println("sql", sql)
	results, err := db2.Query(sql)
	if err != nil {
		msg = err.Error()
		log.Errorf("err:%s ", msg)
		c.JSON(http.StatusOK, gin.H{"success": false, "code": 200, "msg": msg})
		return
	}

	fmt.Println("bb", values, fields, results)
	c.JSON(http.StatusOK, gin.H{"success": true, "code": 200, "msg": msg})
}

// 获取所有记录列表
/*
未做的：
1.分页
2.条件查询
3.统计
4.分组
*/
func GetAllClasses(c *gin.Context) {
	db2 := models.GetDB()
	tableName := c.Param("tableName")

	// limit := c.Query("limit")
	// offset := c.Query("offset")

	limit := 2
	offset := 1

	fmt.Println("limit", limit, offset)

	type result map[string]*string
	var info []result

	db2.Table(tableName).AllCols().Limit(limit, offset).Find(&info)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": info})
}
