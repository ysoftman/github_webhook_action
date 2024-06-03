package github_webhook_action

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine/v2"
	appenginelog "google.golang.org/appengine/v2/log"
)

type GAERouter struct {
	gwh       *GithubWebhook
	allowCORS bool
}

func NewGAERouter(gwh *GithubWebhook, allowCORS bool) *GAERouter {
	return &GAERouter{
		gwh:       gwh,
		allowCORS: allowCORS,
	}
}

func (gae *GAERouter) Start() {
	http.HandleFunc("/", gae.handlerIndex)
	http.HandleFunc("/v1/version", gae.handlerVersion)
	http.HandleFunc("/v1/log", gae.handlerLog)
	http.HandleFunc("/v1/webhook", gae.handlerWebhook)
	appengine.Main()
}

func (gae *GAERouter) SetCommonResponseHeader(w http.ResponseWriter) {
	if gae.allowCORS {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "get")
	}
}

func (gae *GAERouter) handlerIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/ 요청 처리")
	out := `# github webhook action
github webhook 을 받아 필요한 액션을 테스트하는 하는 서버입니다.

# APIs
https://github-webhook-action.appspot.com/v1/version
https://github-webhook-action.appspot.com/v1/webhook

# github
https://github.com/ysoftman/github_webhook_action
`
	gae.SetCommonResponseHeader(w)
	fmt.Fprintln(w, out)
}

func (gae *GAERouter) handlerVersion(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/version 요청 처리")
	gae.SetCommonResponseHeader(w)
	fmt.Fprintln(w, Conf.BuildTime)
}

func (gae *GAERouter) handlerLog(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/log 요청 처리")
	gae.SetCommonResponseHeader(w)
	fmt.Fprintln(w, TailLog())
}

func (gae *GAERouter) handlerWebhook(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/webhook 요청 처리")
	gae.SetCommonResponseHeader(w)
	gae.gwh.githubWebhook(r)
}
