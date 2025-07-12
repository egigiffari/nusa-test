package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HttpServer(routerGroup func(*gin.RouterGroup)) *http.Server {
	return NewHttpServer(":"+os.Getenv("PORT"), routerGroup)
}

func NewHttpServer(addr string, routerGroup func(*gin.RouterGroup)) *http.Server {
	mode := os.Getenv("DEBUG_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	routerGroup(engine.Group("/api"))
	return &http.Server{
		Addr:    addr,
		Handler: engine.Handler(),
	}
}
