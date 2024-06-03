package github_webhook_action

import "github.com/BurntSushi/toml"

var Conf ConfigTOML

type ConfigTOML struct {
	Name      string `toml:"Name"`
	BuildTime string `toml:"BuildTime"`
	Server    struct {
		Port             int    `toml:"Port"`
		LogLevel         string `toml:"LogLevel"`
		LogIsJSONFormat  bool   `toml:"LogIsJsonFormat"`
		WebhookSecretKey string `toml:"WebhookSecretKey"`
	} `toml:"server"`
	Hook []struct {
		RepoName string `toml:"RepoName"`
	} `toml:"hook"`
	Action struct {
		Enable bool   `toml:"Enable"`
		Method string `toml:"Method"`
		Auth   string `toml:"Auth"`
		URL    string `toml:"URL"`
	} `toml:"action"`
}

func LoadConfig() {
	toml.DecodeFile("config.toml", &Conf)
}
