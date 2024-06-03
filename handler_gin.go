package github_webhook_action

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/v2"
	appenginelog "google.golang.org/appengine/v2/log"
)

type GinRouter struct {
	gwh       *GithubWebhook
	allowCORS bool
}

func NewGinRouter(gwh *GithubWebhook, allowCORS bool) *GinRouter {
	return &GinRouter{
		gwh:       gwh,
		allowCORS: allowCORS,
	}
}

func (gr *GinRouter) Start() {
	r := gin.New()
	r.GET("/", gr.ginHandlerIndex)
	v1 := r.Group("/v1")
	{
		v1.GET("/version", gr.ginHandlerVersion)
		v1.GET("/test-sendmessage", gr.ginHandlerTestSendMessage)
		v1.GET("/log", gr.ginHandlerLog)
		v1.POST("/webhook", gr.ginHandlerGithubWebhook)
	}
	r.RedirectTrailingSlash = false
	//http.Handle("/", r)
	//appengine.Main()
	r.Run(fmt.Sprintf(":%v", Conf.Server.Port))
}

func (gr *GinRouter) SetCommonResponseHeader(gc *gin.Context) {
	if gr.allowCORS {
		gc.Header("Access-Control-Allow-Origin", "*")
		gc.Header("Access-Control-Allow-Methods", "get")
	}
}

func (gr *GinRouter) ginHandlerIndex(gc *gin.Context) {
	if appengine.IsAppEngine() {
		ctx := appengine.NewContext(gc.Request)
		appenginelog.Infof(ctx, "/ 요청 처리")
	}
	Zerologger.Info().Str("handler", "ginHandlerIndex").Msg("requst")
	out := `# github webhook action
github Webhook 을 받아 필요한 액션을 테스트하는 하는 서버입니다.

# APIs
https://github-webhook-action.appspot.com/v1/version
https://github-webhook-action.appspot.com/v1/test-sendmessage
https://github-webhook-action.appspot.com/v1/log
https://github-webhook-action.appspot.com/v1/webhook

# github
https://github.com/ysoftman/github_webhook_action
`
	gr.SetCommonResponseHeader(gc)
	gc.String(http.StatusOK, out)
}

func (gr *GinRouter) ginHandlerTestSendMessage(gc *gin.Context) {
	if appengine.IsAppEngine() {
		ctx := appengine.NewContext(gc.Request)
		appenginelog.Infof(ctx, "/test-sendmessage 요청 처리")
	}
	Zerologger.Info().Str("handler", "ginHandlerTestSendMessage").Msg("requst")
	gc.JSON(http.StatusOK, "test send message(check server log)")
	gr.SetCommonResponseHeader(gc)
	gr.gwh.sender.SendMessage("this is test message to check sendmessage")
}

func (gr *GinRouter) ginHandlerVersion(gc *gin.Context) {
	if appengine.IsAppEngine() {
		ctx := appengine.NewContext(gc.Request)
		appenginelog.Infof(ctx, "/version 요청 처리")
	}
	Zerologger.Info().Str("handler", "ginHandlerVersion").Msg("requst")
	gr.SetCommonResponseHeader(gc)
	gc.JSON(http.StatusOK, Conf.BuildTime)
}
func (gr *GinRouter) ginHandlerLog(gc *gin.Context) {
	if appengine.IsAppEngine() {
		ctx := appengine.NewContext(gc.Request)
		appenginelog.Infof(ctx, "/log 요청 처리")
	}
	Zerologger.Info().Str("handler", "ginHandlerVersion").Msg("requst")
	gr.SetCommonResponseHeader(gc)
	gc.JSON(http.StatusOK, TailLog())
}

func (gr *GinRouter) ginHandlerGithubWebhook(gc *gin.Context) {
	if appengine.IsAppEngine() {
		ctx := appengine.NewContext(gc.Request)
		appenginelog.Infof(ctx, "/webhook 요청 처리")
	}
	Zerologger.Info().Str("handler", "ginHandlerGithubWebhook").Msg("requst")
	gr.SetCommonResponseHeader(gc)
	gr.gwh.githubWebhook(gc.Request)
}
