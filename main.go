package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	serverType := flag.String("servertype", "gae", "noraml|gae(google app engin)")
	flag.Parse()
	log.Println("servertype :", *serverType)

	//sender := NewSender() // use default sender
	sender := &mySender{} // use custom sender
	gwh := NewGithubWebhook(sender)
	if *serverType == "normal" {
		// 일반 서버 환경으로 운영시
		NewGinRouter(gwh).Start()
	} else if *serverType == "gae" {
		// GAE(google app engine) 환경으로 운영시
		NewGAERouter(gwh).Start()
	}
	fmt.Println("wrong servertype")
}

type mySender struct {
}

func (s *mySender) SendMessage(msg string) {
	zerologger.Info().Msgf("[my SendMessage] msg:%v", msg)
}
