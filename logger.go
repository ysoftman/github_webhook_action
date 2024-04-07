package main

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

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
