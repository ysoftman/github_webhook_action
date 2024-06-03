package main

import (
	"strings"

	"github.com/go-resty/resty/v2"
	gwa "github.com/ysoftman/github_webhook_action"
)

func main() {
	//sender := gwa.NewSender() // use default sender
	sender := &mySender{} // use custom sender
	gwh := gwa.NewGithubWebhook(sender)
	gwa.NewGAERouter(gwh, true).Start()
}

type mySender struct {
}

func (s *mySender) SendMessage(msg string) {
	gwa.Zerologger.Info().Msgf("[my SendMessage] msg:%v", msg)
	if !gwa.Conf.Action.Enable {
		gwa.Zerologger.Info().Msg("Action is disabled")
		return
	}

	reqBody := struct {
		Message string `json:"msg"`
	}{Message: msg}
	isValidHook := false
	for _, v := range gwa.Conf.Hook {
		if strings.Contains(msg, v.RepoName) {
			isValidHook = true
			gwa.Zerologger.Info().Str("hook from:", v.RepoName).Msg("")
			break
		}
	}
	if !isValidHook {
		gwa.Zerologger.Error().Msg("[my SendMessage] can't find repo name in msg")
	}

	client := resty.New()
	req := client.R().SetHeader("Accept", "application/json").SetBody(&reqBody)
	if len(gwa.Conf.Action.Auth) > 0 {
		req = req.SetAuthToken(gwa.Conf.Action.Auth)
	}

	var resp *resty.Response
	var err error
	if strings.ToLower(gwa.Conf.Action.Method) == "post" {
		resp, err = req.Post(gwa.Conf.Action.URL)
	} else if strings.ToLower(gwa.Conf.Action.Method) == "get" {
		resp, err = req.SetQueryParams(map[string]string{
			"param1": "apple",
			"param2": "lemon"}).Get(gwa.Conf.Action.URL)
	}
	if err != nil {
		gwa.Zerologger.Error().Err(err).Msg("[my SendMessage] failed to sendMessage")
	}
	gwa.Zerologger.Info().Msgf("[my SendMessage] resp:%v", resp)
}
