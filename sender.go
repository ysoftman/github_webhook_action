package main

import (
	"strings"

	"github.com/go-resty/resty/v2"
)

type SenderInterface interface {
	SendMessage(msg string)
}

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) SendMessage(msg string) {
	zerologger.Info().Msgf("msg:%v", msg)
	if !conf.Action.API.Enable {
		zerologger.Info().Msg("action api is disabled")
		return
	}
	client := resty.New()
	reqBody := struct {
		Field1  string `json:"field1"`
		Message string `json:"msg"`
	}{
		"aaa",
		msg,
	}
	req := client.R().SetHeader("Accept", "application/json").SetBody(&reqBody)
	if len(conf.Action.API.Auth) > 0 {
		req = req.SetAuthToken(conf.Action.API.Auth)
	}

	var resp *resty.Response
	var err error
	if strings.ToLower(conf.Action.API.Mothod) == "post" {
		resp, err = req.Post(conf.Action.API.URL)
	} else if strings.ToLower(conf.Action.API.Mothod) == "get" {
		resp, err = req.SetQueryParams(map[string]string{
			"param1": "apple",
			"param2": "lemon"}).Get(conf.Action.API.URL)
	}
	if err != nil {
		zerologger.Error().Err(err).Msg("failed to sendMessage")
	}
	zerologger.Info().Msgf("resp:%v", resp)
}
