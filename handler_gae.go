package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine/v2"
	appenginelog "google.golang.org/appengine/v2/log"
)

type GAERouter struct {
}

func NewGAERouter() *GAERouter {
	return &GAERouter{}
}

func (gae *GAERouter) Start() {
	http.HandleFunc("/", gae.handlerIndex)
	http.HandleFunc("/v1/version", gae.handlerVersion)
	http.HandleFunc("/v1/webhook", gae.handlerWebhook)
	appengine.Main()
}

func (gae *GAERouter) handlerIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/ 요청 처리")
	out := `
# github webhook action
github webhook 을 받아 필요한 액션을 테스트하는 하는 서버입니다.

# app engine 으로 운영
https://github-webhook-action.appspot.com/v1/webhook/

# github
https://github.com/ysoftman/github_webhook_action
`
	fmt.Fprintln(w, out)
}

func (gae *GAERouter) handlerVersion(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/version 요청 처리")
	fmt.Fprintln(w, buildtime)
}

func (gae *GAERouter) handlerWebhook(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/webhook 요청 처리")
	githubWebhook(r)
}
