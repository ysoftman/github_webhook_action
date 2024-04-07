package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog"
)

type configTOML struct {
	Name   string `toml:"Name"`
	Server struct {
		Port             int    `toml:"Port"`
		LogLevel         string `toml:"LogLevel"`
		LogIsJsonFormat  bool   `toml:"LogIsJsonFormat"`
		WebhookSecretKey string `toml:"WebhookSecretKey"`
	} `toml:"server"`
	Action struct {
		API struct {
			Enable      bool   `toml:"Enable"`
			URL         string `toml:"URL"`
			Auth        string `toml:"Auth"`
			Mothod      string `toml:"Mothod"`
			RequestBody string `toml:"RequestBody"`
		} `toml:"api"`
	} `toml:"action"`
}

var buildtime string
var conf configTOML
var logger zerolog.Logger

func main() {
	serverType := flag.String("servertype", "gae", "noraml|gae(google app engin)")
	flag.Parse()
	log.Println("servertype :", *serverType)

	toml.DecodeFile("config.toml", &conf)
	CreateLogger(conf.Server.LogLevel, conf.Server.LogIsJsonFormat)
	logger.Info().Msg("github_webhook_action")
	if *serverType == "normal" {
		gr := NewGinRouter()
		// 일반 서버 환경으로 운영시
		gr.Start()
	} else if *serverType == "gae" {
		// GAE(google app engine) 환경으로 운영시
		gae := NewGAERouter()
		gae.Start()
	}
	fmt.Println("wrong servertype")
}
