package main

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var zerologger zerolog.Logger

func CreateLogger(logLevelString string, isJson bool) {
	logLevel, _ := zerolog.ParseLevel(logLevelString)
	var writer io.Writer
	if isJson {
		writer = os.Stdout
	} else {
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	}
	zerologger = zerolog.New(writer).With().Timestamp().Logger().Level(logLevel)
}
