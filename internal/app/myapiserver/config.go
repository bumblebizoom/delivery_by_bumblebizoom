package myapiserver

import "github.com/bumblebizoom/delivery_by_bumblebizoom/internal/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
