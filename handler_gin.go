package github_webhook_action

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	v1 := r.Group("/v1")
	{
		v1.GET("/test-sendmessage", gr.ginHandlerTestSendMessage)
		v1.GET("/version", gr.ginHandlerVersion)
		v1.GET("/log", gr.ginHandlerLog)
		v1.POST("/webhook", gr.ginHandlerGithubWebhook)
	}
	r.RedirectTrailingSlash = false
	r.Run(fmt.Sprintf(":%v", Conf.Server.Port))
}

func (gr *GinRouter) SetCommonResponseHeader(gc *gin.Context) {
	if gr.allowCORS {
		gc.Header("Access-Control-Allow-Origin", "*")
		gc.Header("Access-Control-Allow-Methods", "get")
	}
}

func (gr *GinRouter) ginHandlerTestSendMessage(gc *gin.Context) {
	Zerologger.Info().Str("handler", "ginHandlerTestSendMessage").Msg("requst")
	gc.JSON(http.StatusOK, "test send message(check server log)")
	gr.SetCommonResponseHeader(gc)
	gr.gwh.sender.SendMessage("this is test message to check sendmessage")
}

func (gr *GinRouter) ginHandlerVersion(gc *gin.Context) {
	Zerologger.Info().Str("handler", "ginHandlerVersion").Msg("requst")
	gr.SetCommonResponseHeader(gc)
	gc.JSON(http.StatusOK, Conf.BuildTime)
}
func (gr *GinRouter) ginHandlerLog(gc *gin.Context) {
	Zerologger.Info().Str("handler", "ginHandlerVersion").Msg("requst")
	gr.SetCommonResponseHeader(gc)
	gc.JSON(http.StatusOK, TailLog())
}

func (gr *GinRouter) ginHandlerGithubWebhook(gc *gin.Context) {
	Zerologger.Info().Str("handler", "ginHandlerGithubWebhook").Msg("requst")
	gr.SetCommonResponseHeader(gc)
	gr.gwh.githubWebhook(gc.Request)
}
