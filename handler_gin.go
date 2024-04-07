package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
}

func NewGinRouter() *GinRouter {
	return &GinRouter{}
}

func (gr *GinRouter) Start() {
	r := gin.New()
	v1 := r.Group("/v1")
	{
		v1.GET("/version", gr.ginHandlerVersion)
		v1.POST("/webhook", gr.ginHandlerGithubWebhook)
	}
	r.Run(fmt.Sprintf(":%v", conf.Server.Port))
}

func (gr *GinRouter) ginHandlerVersion(gc *gin.Context) {
	logger.Info().Str("handler", "ginHandlerVersion").Msg("requst")
	gc.JSON(http.StatusOK, buildtime)
}

func (gr *GinRouter) ginHandlerGithubWebhook(gc *gin.Context) {
	logger.Info().Str("handler", "ginHandlerGithubWebhook").Msg("requst")
	githubWebhook(gc.Request)
}
