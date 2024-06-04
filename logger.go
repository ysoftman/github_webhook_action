package main

import (
	"io"
	"os"

	"github.com/nxadm/tail"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Zerologger zerolog.Logger

func CreateLogger(logFile, logLevelString string, isJson bool) {
	logLevel, _ := zerolog.ParseLevel(logLevelString)
	var writers []io.Writer
	if isJson {
		writers = append(writers, os.Stdout)
	} else {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout})
	}
	lj := &lumberjack.Logger{
		// 기본 로그 파일 명
		Filename: logFile,
		// 로그 파일당 최대 허용 크기(megabytes) - rotate 조건
		// MaxSize is the maximum size in megabytes of the log file before it gets
		// rotated. It defaults to 100 megabytes.
		// MaxSize 보다 커야만 rotate 된다.
		// MaxSize: 999999,
		MaxSize: 1,

		// old 로그 유지 조건 - MaxAge or MaxBackups
		// old 로그 유지 기간(days)
		// MaxAge is the maximum number of days to retain old log files based on the
		// timestamp encoded in their filename.  Note that a day is defined as 24
		// hours and may not exactly correspond to calendar days due to daylight
		// savings, leap seconds, etc. The default is not to remove old log files
		// based on age.
		MaxAge: 365 * 100,

		// old 로그 유지 개수
		// MaxBackups is the maximum number of old log files to retain.  The default
		// is to retain all old log files (though MaxAge may still cause them to get
		// deleted.)
		MaxBackups: 2,

		// 로컬 시간으로 파일명(타임스탬프)사용, 기본 UTC
		// LocalTime determines if the time used for formatting the timestamps in
		// backup files is the computer's local time.  The default is to use UTC
		// time.
		LocalTime: true,

		// 압축여부
		// Compress determines if the rotated log files should be compressed
		// using gzip. The default is not to perform compression.
		Compress: false,
	}
	writers = append(writers, lj)
	mw := io.MultiWriter(writers...)
	Zerologger = zerolog.New(mw).With().Timestamp().Logger().Level(logLevel)
}

func TailLog() string {
	finfo, err := os.Stat(Conf.Server.LogFile)
	if err != nil {
		return ""
	}
	var offset int64 = 5000
	if offset > finfo.Size() {
		offset = finfo.Size()
	}
	t, err := tail.TailFile(Conf.Server.LogFile,
		tail.Config{
			Location: &tail.SeekInfo{Offset: -offset, Whence: io.SeekEnd},
			ReOpen:   false, // cannot set ReOpen without Follow.
			Follow:   false,
		})
	if err != nil {
		return ""
	}
	log := ""
	for line := range t.Lines {
		log += line.Text
		log += "\n"
	}
	return log
}
