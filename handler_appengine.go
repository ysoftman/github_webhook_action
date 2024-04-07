package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine/v2"
	appenginelog "google.golang.org/appengine/v2/log"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
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
func handlerVersion(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/version 요청 처리")
	fmt.Fprintln(w, buildtime)
}
func handlerWebhook(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/webhook 요청 처리")
	githubWebhook(r)
}
