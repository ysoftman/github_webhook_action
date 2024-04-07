package github_webhook_action

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	gwh *GithubWebhook
}

func NewGinRouter(gwh *GithubWebhook) *GinRouter {
	return &GinRouter{gwh: gwh}
}

func (gr *GinRouter) Start() {
	r := gin.New()
	v1 := r.Group("/v1")
	{
		v1.GET("/test_sendmessage", gr.ginHandlerTestSendMessage)
		v1.GET("/version", gr.ginHandlerVersion)
		v1.POST("/webhook", gr.ginHandlerGithubWebhook)
	}
	r.Run(fmt.Sprintf(":%v", conf.Server.Port))
}

func (gr *GinRouter) ginHandlerTestSendMessage(gc *gin.Context) {
	zerologger.Info().Str("handler", "ginHandlerTestSendMessage").Msg("requst")
	gc.JSON(http.StatusOK, "test send message(check server log)")
	gr.gwh.sender.SendMessage("this is test message to check sendmessage")
}

func (gr *GinRouter) ginHandlerVersion(gc *gin.Context) {
	zerologger.Info().Str("handler", "ginHandlerVersion").Msg("requst")
	gc.JSON(http.StatusOK, buildtime)
}

func (gr *GinRouter) ginHandlerGithubWebhook(gc *gin.Context) {
	zerologger.Info().Str("handler", "ginHandlerGithubWebhook").Msg("requst")
	gr.gwh.githubWebhook(gc.Request)
}
