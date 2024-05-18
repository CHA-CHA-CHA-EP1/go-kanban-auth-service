package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/CHA-CHA-CHA-EP1/go-kanban-auth-service/config"
	routes "github.com/CHA-CHA-CHA-EP1/go-kanban-auth-service/pkg"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	cfg := config.InitialConfig()
	var (
		logger *zap.Logger
		err    error
	)

	if cfg.Env == "dev" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	logger.Info("config loaded")

	router := routes.NewRouter().RegisterRoutes()

	go Run(router, cfg.Server.Port)

	GracefulShutdown(router)
	logger.Info("shutting down server...")
}

func Run(router *echo.Echo, port int) {
	router.Logger.Fatal(router.Start(":" + strconv.Itoa(port)))
}

func GracefulShutdown(server *echo.Echo) {

	shutdownContext, cencelShutdown := context.WithTimeout(context.Background(), 10*time.Second)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	defer cencelShutdown()

	if err := server.Shutdown(shutdownContext); err != nil {
		server.Logger.Fatal(err)
	}
}
