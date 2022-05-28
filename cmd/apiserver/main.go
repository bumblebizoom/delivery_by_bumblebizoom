package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	logger "github.com/bumblebizoom/bumblogger"
	"github.com/bumblebizoom/delivery_by_bumblebizoom/internal/app/myapiserver"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/appserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := myapiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logger.Fatal("Config file error", err)
	}

	s := myapiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
