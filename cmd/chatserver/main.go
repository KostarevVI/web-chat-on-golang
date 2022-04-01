package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"web-chat-on-golang/internal/app/chatserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/chatserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := chatserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := chatserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
