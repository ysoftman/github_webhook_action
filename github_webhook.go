package github_webhook_action

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
)

type GithubWebhook struct {
	sender SenderInterface
}

func NewGithubWebhook(sender SenderInterface) *GithubWebhook {
	LoadConfig()
	CreateLogger(Conf.Server.LogLevel, Conf.Server.LogIsJSONFormat)
	return &GithubWebhook{sender: sender}
}

func (gwh *GithubWebhook) githubWebhook(req *http.Request) {
	payload, err1 := github.ValidatePayload(req, []byte(Conf.Server.WebhookSecretKey))
	if err1 != nil {
		Zerologger.Error().Err(err1).Msg("failed ValidatePayload")
		return
	}
	event, err2 := github.ParseWebHook(github.WebHookType(req), payload)
	if err2 != nil {
		Zerologger.Error().Err(err2).Msg("failed ParseWebHook")
		return
	}
	webhookType := github.WebHookType(req)
	Zerologger.Info().Msgf("github WebHookType:%v", webhookType)
	switch event := event.(type) {
	case *github.CommitCommentEvent:
		gwh.githubCommitCommentEvent(event)
	case *github.PushEvent:
		gwh.githubPushEvent(event)
	case *github.PullRequestEvent:
		gwh.githubPullRequestEvent(event)
	case *github.PullRequestReviewEvent:
		gwh.githubPullRequestReviewEvent(event)
	case *github.PullRequestReviewCommentEvent:
		gwh.githubPullRequestReviewCommentEvent(event)
	default:
		Zerologger.Info().Msgf("github WebHookType:%v", webhookType)
	}
}
func (gwh *GithubWebhook) githubCommitCommentEvent(event *github.CommitCommentEvent) {
	msg := fmt.Sprintf("[CommitComment-%v] sender:%v(%v) comment:%v link:%v",
		event.GetAction(),
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.GetComment(),
		event.Comment.HTMLURL)
	gwh.sender.SendMessage(msg)
}
func (gwh *GithubWebhook) githubPushEvent(event *github.PushEvent) {
	msg := fmt.Sprintf("[Push-%v] sender:%v(%v) pusher:%v link:%v",
		*event.HeadCommit.Message,
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.Pusher.GetName(),
		event.Repo.GetHTMLURL())
	gwh.sender.SendMessage(msg)
}
func (gwh *GithubWebhook) githubPullRequestEvent(event *github.PullRequestEvent) {
	msg := fmt.Sprintf("[PullRequest-%v] sender:%v(%v) number:%v link:%v",
		event.GetAction(),
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.GetNumber(),
		*event.PullRequest.HTMLURL)
	gwh.sender.SendMessage(msg)
}
func (gwh *GithubWebhook) githubPullRequestReviewEvent(event *github.PullRequestReviewEvent) {
	msg := fmt.Sprintf("[PullRequestReview-%v] sender:%v(%v) review:%v link:%v",
		event.GetAction(),
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.GetReview().String(),
		event.GetReview().GetHTMLURL())
	gwh.sender.SendMessage(msg)
}
func (gwh *GithubWebhook) githubPullRequestReviewCommentEvent(event *github.PullRequestReviewCommentEvent) {
	msg := fmt.Sprintf("[PullRequestReviewComment-%v] sender:%v(%v) comment:%v link:%v",
		event.GetAction(),
		event.Sender.GetLogin(),
		event.Sender.GetName(),
		event.Comment.String(),
		event.GetComment().GetURL())
	gwh.sender.SendMessage(msg)
}
