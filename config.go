package github_webhook_action

import "github.com/BurntSushi/toml"

var conf ConfigTOML

type ConfigTOML struct {
	Name      string `toml:"Name"`
	BuildTime string `toml:"BuildTime"`
	Server    struct {
		Port             int    `toml:"Port"`
		LogLevel         string `toml:"LogLevel"`
		LogIsJSONFormat  bool   `toml:"LogIsJsonFormat"`
		WebhookSecretKey string `toml:"WebhookSecretKey"`
	} `toml:"server"`
	Action struct {
		API struct {
			Enable bool   `toml:"Enable"`
			Mothod string `toml:"Mothod"`
			Auth   string `toml:"Auth"`
			URL    string `toml:"URL"`
		} `toml:"api"`
		Target []struct {
			RepoName    string `toml:"RepoName"`
			TargetID    int    `toml:"TargetID"`
			RequestBody string `toml:"RequestBody"`
		} `toml:"target"`
	} `toml:"action"`
}

func LoadConfig() {
	toml.DecodeFile("config.toml", &conf)
}
