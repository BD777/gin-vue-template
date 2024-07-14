package service

import (
	"fmt"
	"gin-vue-template/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service struct {
	HTTPServer *http.Server
}

func NewHTTPServer(engine *gin.Engine, cfg *config.Config) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.HTTPServer.Port),
		Handler:        engine,
		ReadTimeout:    time.Duration(cfg.HTTPServer.ReadTimeout) * time.Millisecond,
		WriteTimeout:   time.Duration(cfg.HTTPServer.WriteTimeout) * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}
}
