package models

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/luxingwen/Bass/config"
)

func GetDB() *xorm.Engine {

	engine, err := xorm.NewEngine("mysql", "root:@/"+config.MySqlConf.DbName+"?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	engine.ShowSQL(true)
	engine.Logger().SetLevel(0)

	f, err := os.Create("log/sql.log")
	if err != nil {
		println(err.Error())

	}
	engine.SetLogger(xorm.NewSimpleLogger(f))

	return engine
}
