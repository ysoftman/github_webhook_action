package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginHandlerVersion(gc *gin.Context) {
	logger.Info().Str("handler", "ginHandlerVersion").Msg("requst")
	gc.JSON(http.StatusOK, buildtime)
}
func ginHandlerGithubWebhook(gc *gin.Context) {
	logger.Info().Str("handler", "ginHandlerGithubWebhook").Msg("requst")
	githubWebhook(gc.Request)
}
