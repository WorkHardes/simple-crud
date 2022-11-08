package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/simple-crud/internal/config"
	"example.com/simple-crud/internal/routers"
	"example.com/simple-crud/internal/server"
	logs "example.com/simple-crud/pkg/logger"
)

func main() {
	// Init logger
	logger := logs.NewCustomLogger()

	// Get server config
	cfg, err := config.Init()
	if err != nil {
		logger.Error("error when init config. detail: %v", err)
		os.Exit(1)
	}

	srvHandler := routers.Init()

	// Make server with config
	srv := server.NewServer(cfg, srvHandler)

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server. detail: %s", err.Error())
		}
	}()

	logger.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server. detail: %v", err)
	}
}
