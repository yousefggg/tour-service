package config

import (
	"github.com/joho/godotenv"
	commonconfig "github.com/yousefggg/common-lib/pkg/config"
	"github.com/yousefggg/common-lib/pkg/logger"
)

type Config struct {
	*commonconfig.Config
}

func LoadConfig() *Config {

	_ = godotenv.Load()

	cfg, err := commonconfig.LoadConfig()
	if err != nil {
		logger.Error("failed to load common config", "error", err)
		panic(err)
	}

	return &Config{
		Config: cfg,
	}
}