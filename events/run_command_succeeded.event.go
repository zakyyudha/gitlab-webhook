package events

import (
	"fmt"
	"gitlab-webhook/config"
	"gitlab-webhook/dto"
	"gitlab-webhook/internal/libraries/discord"
	discordDto "gitlab-webhook/internal/libraries/discord/dto"
	"gitlab-webhook/internal/libraries/event"
)

// RunCommandSucceeded ..
const RunCommandSucceeded event.Name = "RunCommandSucceeded"

// RunCommandSucceededEvent ..
type RunCommandSucceededEvent struct {
	Config  *config.GitlabConfig
	Request *dto.GitlabWebhooks
}

func (e RunCommandSucceededEvent) SendToDiscord() {
	webhookReq := discordDto.WebhookRequest{}
	content := fmt.Sprintf("Deploy succeded on `%v`\n", e.Config.Appname)
	webhookReq.Content = content

	discord.Send(webhookReq, nil, e.Config.DiscordWebhook)
}