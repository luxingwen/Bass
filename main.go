package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/cihub/seelog"
	"github.com/luxingwen/Bass/config"

	"github.com/luxingwen/Bass/router"
)

func main() {
	r := router.Router()
	s := &http.Server{
		Addr:           ":" + config.ServerConf.Port,
		Handler:        r,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 日志配置
	defer log.Flush()
	//加载配置文件
	logger, err := log.LoggerFromConfigAsFile("conf/seelog.xml")

	if err != nil {
		fmt.Println("parse seelog.xml error")
	}
	//替换记录器
	log.ReplaceLogger(logger)

	// log.Info("Hello world from Seelog!")

	s.ListenAndServe()
}
