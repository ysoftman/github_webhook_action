package main

import (
	"strings"

	"github.com/go-resty/resty/v2"
)

func main() {
	//sender := NewSender() // use default sender
	sender := &mySender{} // use custom sender
	gwh := NewGithubWebhook(sender)
	NewGAERouter(gwh, true).Start()
}

type mySender struct {
}

func (s *mySender) SendMessage(msg string) {
	Zerologger.Info().Msgf("[my SendMessage] msg:%v", msg)
	if !Conf.Action.Enable {
		Zerologger.Info().Msg("Action is disabled")
		return
	}

	reqBody := struct {
		Message string `json:"msg"`
	}{Message: msg}
	isValidHook := false
	for _, v := range Conf.Hook {
		if strings.Contains(msg, v.RepoName) {
			isValidHook = true
			Zerologger.Info().Str("hook from:", v.RepoName).Msg("")
			break
		}
	}
	if !isValidHook {
		Zerologger.Error().Msg("[my SendMessage] can't find repo name in msg")
	}

	client := resty.New()
	req := client.R().SetHeader("Accept", "application/json").SetBody(&reqBody)
	if len(Conf.Action.Auth) > 0 {
		req = req.SetAuthToken(Conf.Action.Auth)
	}

	var resp *resty.Response
	var err error
	if strings.ToLower(Conf.Action.Method) == "post" {
		resp, err = req.Post(Conf.Action.URL)
	} else if strings.ToLower(Conf.Action.Method) == "get" {
		resp, err = req.SetQueryParams(map[string]string{
			"param1": "apple",
			"param2": "lemon"}).Get(Conf.Action.URL)
	}
	if err != nil {
		Zerologger.Error().Err(err).Msg("[my SendMessage] failed to sendMessage")
	}
	Zerologger.Info().Msgf("[my SendMessage] resp:%v", resp)
}
