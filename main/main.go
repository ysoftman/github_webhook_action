package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
	gwa "github.com/ysoftman/github_webhook_action"
)

func main() {
	serverType := flag.String("servertype", "gae", "noraml|gae(google app engin)")
	flag.Parse()
	log.Println("servertype :", *serverType)

	//sender := gwa.NewSender() // use default sender
	sender := &mySender{} // use custom sender
	gwh := gwa.NewGithubWebhook(sender)
	if *serverType == "normal" {
		// 일반 서버 환경으로 운영시
		gwa.NewGinRouter(gwh).Start()
	} else if *serverType == "gae" {
		// GAE(google app engine) 환경으로 운영시
		gwa.NewGAERouter(gwh).Start()
	}
	fmt.Println("wrong servertype")
}

type mySender struct {
}

func (s *mySender) SendMessage(msg string) {
	gwa.Zerologger.Info().Msgf("[my SendMessage] msg:%v", msg)
	if !gwa.Conf.Action.API.Enable {
		gwa.Zerologger.Info().Msg("action api is disabled")
		return
	}

	reqBody := struct {
		To      int    `json:"to"`
		Message string `json:"msg"`
	}{Message: msg}
	for _, v := range gwa.Conf.Action.Target {
		if strings.Contains(msg, v.RepoName) {
			reqBody.To = v.TargetID
			gwa.Zerologger.Info().Int("target ID", v.TargetID).Msg("")
			break
		}
	}

	client := resty.New()
	req := client.R().SetHeader("Accept", "application/json").SetBody(&reqBody)
	if len(gwa.Conf.Action.API.Auth) > 0 {
		req = req.SetAuthToken(gwa.Conf.Action.API.Auth)
	}

	var resp *resty.Response
	var err error
	if strings.ToLower(gwa.Conf.Action.API.Mothod) == "post" {
		resp, err = req.Post(gwa.Conf.Action.API.URL)
	} else if strings.ToLower(gwa.Conf.Action.API.Mothod) == "get" {
		resp, err = req.SetQueryParams(map[string]string{
			"param1": "apple",
			"param2": "lemon"}).Get(gwa.Conf.Action.API.URL)
	}
	if err != nil {
		gwa.Zerologger.Error().Err(err).Msg("[my SendMessage] failed to sendMessage")
	}
	gwa.Zerologger.Info().Msgf("[my SendMessage] resp:%v", resp)
}
