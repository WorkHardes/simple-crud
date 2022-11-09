package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	SrvCfg ServerConfig
}

type ServerConfig struct {
	Host           string
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func getEnv(envName string) (string, error) {
	envVar := os.Getenv(envName)
	if envVar == "" {
		return "", fmt.Errorf("%s is nil!", envName)
	}

	return envVar, nil
}

func initServerConfig() (*ServerConfig, error) {
	var srvCfg ServerConfig

	var err error

	srvCfg.Host, err = getEnv("SERVER_HOST")
	if err != nil {
		return nil, fmt.Errorf("error when getting server host. detail: %v", err)
	}

	srvCfg.Port, err = getEnv("SERVER_PORT")
	if err != nil {
		return nil, fmt.Errorf("error when getting server port. detail: %v", err)
	}

	readTimeoutStr, err := getEnv("SERVER_READ_TIMEOUT")
	if err != nil {
		return nil, fmt.Errorf("error when getting server read timeout. detail: %v", err)
	}

	readTimeout, err := strconv.Atoi(readTimeoutStr)
	if err != nil {
		return nil, fmt.Errorf("error when converting server read timeout: '%s'."+
			" detail: %v", readTimeoutStr, err)
	}

	srvCfg.ReadTimeout = time.Duration(readTimeout) * time.Second

	writeTimeoutStr, err := getEnv("SERVER_WRITE_TIMEOUT")
	if err != nil {
		return nil, fmt.Errorf("error when getting server read timeout. detail: %v", err)
	}

	writeTimeout, err := strconv.Atoi(writeTimeoutStr)
	if err != nil {
		return nil, fmt.Errorf("error when converting server write timeout: '%s'."+
			" detail: %v", writeTimeoutStr, err)
	}

	srvCfg.ReadTimeout = time.Duration(writeTimeout) * time.Second

	maxHeaderBytesStr, err := getEnv("SERVER_MAX_HEADER_BYTES")
	if err != nil {
		return nil, fmt.Errorf("error when getting server max header bytes. detail: %v", err)
	}

	srvCfg.MaxHeaderBytes, err = strconv.Atoi(writeTimeoutStr)
	if err != nil {
		return nil, fmt.Errorf("error when converting server max header bytes: '%s'."+
			" detail: %v", maxHeaderBytesStr, err)
	}

	return &srvCfg, nil
}

func Init() (*Config, error) {
	var cfg Config

	srvCfg, err := initServerConfig()
	if err != nil {
		return nil, fmt.Errorf("error when initing server config. detail: %v", err)
	}

	cfg.SrvCfg = *srvCfg

	return &cfg, nil
}
