package github_webhook_action

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var Zerologger zerolog.Logger

func CreateLogger(logLevelString string, isJson bool) {
	logLevel, _ := zerolog.ParseLevel(logLevelString)
	var writer io.Writer
	if isJson {
		writer = os.Stdout
	} else {
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	}
	Zerologger = zerolog.New(writer).With().Timestamp().Logger().Level(logLevel)
}
