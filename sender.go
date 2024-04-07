package github_webhook_action

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
	Zerologger.Info().Msgf("msg:%v", msg)
	if !Conf.Action.API.Enable {
		Zerologger.Info().Msg("action api is disabled")
		return
	}

	reqBody := struct {
		To      int    `json:"to"`
		Message string `json:"msg"`
	}{Message: msg}
	for _, v := range Conf.Action.Target {
		if strings.Contains(msg, v.RepoName) {
			reqBody.To = v.TargetID
			Zerologger.Info().Int("target ID", v.TargetID).Msg("")
			break
		}
	}

	client := resty.New()
	req := client.R().SetHeader("Accept", "application/json").SetBody(&reqBody)
	if len(Conf.Action.API.Auth) > 0 {
		req = req.SetAuthToken(Conf.Action.API.Auth)
	}

	var resp *resty.Response
	var err error
	if strings.ToLower(Conf.Action.API.Mothod) == "post" {
		resp, err = req.Post(Conf.Action.API.URL)
	} else if strings.ToLower(Conf.Action.API.Mothod) == "get" {
		resp, err = req.SetQueryParams(map[string]string{
			"param1": "apple",
			"param2": "lemon"}).Get(Conf.Action.API.URL)
	}
	if err != nil {
		Zerologger.Error().Err(err).Msg("failed to sendMessage")
	}
	Zerologger.Info().Msgf("resp:%v", resp)
}
