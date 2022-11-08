package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	logs "example.com/simple-crud/pkg/logger"
)

type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func getServerConfig() ServerConfig {
	serverConfig := ServerConfig{}
	serverConfig.Host = os.Getenv("SERVER_HOST")
	serverConfig.Port, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))

	readTimeout, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	serverConfig.ReadTimeout = time.Duration(readTimeout) * time.Second
	writeTimeout, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	serverConfig.ReadTimeout = time.Duration(writeTimeout) * time.Second

	return serverConfig
}

func main() {
	// Get server config
	serverConfig := getServerConfig()
	serverAddr := serverConfig.Host + ":" + strconv.Itoa(serverConfig.Port)

	// Make server with config
	s := &http.Server{
		Addr:           serverAddr,
		ReadTimeout:    serverConfig.ReadTimeout,
		WriteTimeout:   serverConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	logger := logs.NewCustomLogger()
	logger.Infof("server starting at %s\n", serverAddr)

	log.Fatal(s.ListenAndServe())
}
