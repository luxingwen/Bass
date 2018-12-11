package main

import (
	"net/http"
	"time"

	"github.com/luxingwen/Bass/config"
	// log "github.com/sirupsen/logrus"

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
	// log.SetLevel(log.DebugLevel)
	s.ListenAndServe()
}
