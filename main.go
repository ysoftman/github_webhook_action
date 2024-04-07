package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"google.golang.org/appengine/v2"
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

func CreateLogger(logLevelString string, isJson bool) {
	logLevel, _ := zerolog.ParseLevel(logLevelString)
	var writer io.Writer
	if isJson {
		writer = os.Stdout
	} else {
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	}
	logger = zerolog.New(writer).With().Timestamp().Logger().Level(logLevel)
}
func main() {
	serverType := flag.String("servertype", "gae", "noraml|gae(google app engin)")
	flag.Parse()
	log.Println("servertype :", *serverType)
	toml.DecodeFile("config.toml", &conf)
	CreateLogger(conf.Server.LogLevel, conf.Server.LogIsJsonFormat)
	logger.Info().Msg("github_webhook_action")
	if *serverType == "normal" {
		// 일반 서버 환경으로 운영시
		r := gin.New()
		v1 := r.Group("/v1")
		{
			v1.GET("/version", ginHandlerVersion)
			v1.POST("/webhook", ginHandlerGithubWebhook)
		}
		r.Run(fmt.Sprintf(":%v", conf.Server.Port))
	} else if *serverType == "gae" {
		// GAE(google app engine) 환경으로 운영시
		http.HandleFunc("/", handlerIndex)
		http.HandleFunc("/v1/version", handlerVersion)
		http.HandleFunc("/v1/webhook", handlerWebhook)
		appengine.Main()
	}
}
