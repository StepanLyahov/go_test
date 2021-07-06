package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort      = "8080"
	defaultHTTPRWTimeout = 10 * time.Second
)

type (
	Config struct {
		HTTP HTTPConfig
	}

	HTTPConfig struct {
		Host         string
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
)

func New() (*Config, error) {
	setDefaults()

	var cfg Config
	if err := unmarshall(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.readTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("http.writeTimeout", defaultHTTPRWTimeout)
}

func unmarshall(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	return nil
}
