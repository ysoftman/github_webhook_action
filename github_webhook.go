package main

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
)

func githubWebhook(req *http.Request) {
	payload, err1 := github.ValidatePayload(req, []byte(conf.Server.WebhookSecretKey))
	if err1 != nil {
		logger.Error().Err(err1).Msg("failed ValidatePayload")
		return
	}
	event, err2 := github.ParseWebHook(github.WebHookType(req), payload)
	if err2 != nil {
		logger.Error().Err(err2).Msg("failed ParseWebHook")
		return
	}
	webhookType := github.WebHookType(req)
	logger.Info().Msgf("github WebHookType:%v", webhookType)
	switch event := event.(type) {
	case *github.CommitCommentEvent:
		githubCommitCommentEvent(event)
	case *github.PushEvent:
		githubPushEvent(event)
	case *github.PullRequestEvent:
		githubPullRequestEvent(event)
	case *github.PullRequestReviewEvent:
		githubPullRequestReviewEvent(event)
	case *github.PullRequestReviewCommentEvent:
		githubPullRequestReviewCommentEvent(event)
	default:
		logger.Info().Msgf("github WebHookType:%v", webhookType)
	}
}
func githubCommitCommentEvent(event *github.CommitCommentEvent) {
	msg := fmt.Sprintf("[CommitComment-%v] sender:%v(%v) comment:%v link:%v",
		event.GetAction(),
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.GetComment(),
		event.Comment.HTMLURL)
	sendMessage(msg)
}
func githubPushEvent(event *github.PushEvent) {
	msg := fmt.Sprintf("[Push-%v] sender:%v(%v) pusher:%v link:%v",
		*event.HeadCommit.Message,
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.Pusher.GetName(),
		event.Repo.GetHTMLURL())
	sendMessage(msg)
}
func githubPullRequestEvent(event *github.PullRequestEvent) {
	msg := fmt.Sprintf("[PullRequest-%v] sender:%v(%v) number:%v link:%v",
		event.GetAction(),
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.GetNumber(),
		*event.PullRequest.HTMLURL)
	sendMessage(msg)
}
func githubPullRequestReviewEvent(event *github.PullRequestReviewEvent) {
	msg := fmt.Sprintf("[PullRequestReview-%v] sender:%v(%v) review:%v link:%v",
		event.GetAction(),
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.GetReview().String(),
		event.GetReview().GetHTMLURL())
	sendMessage(msg)
}
func githubPullRequestReviewCommentEvent(event *github.PullRequestReviewCommentEvent) {
	msg := fmt.Sprintf("[PullRequestReviewComment-%v] sender:%v(%v) comment:%v link:%v",
		event.GetAction(),
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.Comment.String(),
		event.GetComment().GetURL())
	sendMessage(msg)
}
