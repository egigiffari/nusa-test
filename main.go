package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/egigiffari/nusa-test/app"
	"github.com/egigiffari/nusa-test/ports"
	"github.com/egigiffari/nusa-test/server"
	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Background()
	application := app.NewApplication(ctx)

	srv := server.HttpServer(func(routerGroup *gin.RouterGroup) {
		ports.NewHttpHandlers(routerGroup, application)
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")

}
