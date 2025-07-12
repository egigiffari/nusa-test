package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpServer(routerGroup func(*gin.RouterGroup)) *http.Server {
	return NewHttpServer(":8080", routerGroup)
}

func NewHttpServer(addr string, routerGroup func(*gin.RouterGroup)) *http.Server {
	engine := gin.Default()
	routerGroup(engine.Group("/api"))
	return &http.Server{
		Addr:    addr,
		Handler: engine.Handler(),
	}
}
