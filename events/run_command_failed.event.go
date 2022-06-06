package events

import (
	"fmt"
	"gitlab-webhook/config"
	"gitlab-webhook/dto"
	"gitlab-webhook/internal/libraries/discord"
	discordDto "gitlab-webhook/internal/libraries/discord/dto"
	"gitlab-webhook/internal/libraries/event"
)

// RunCommandFailed ..
const RunCommandFailed event.Name = "RunCommandFailed"

// RunCommandFailedEvent ..
type RunCommandFailedEvent struct {
	Config  *config.GitlabConfig
	Request *dto.GitlabWebhooks
	Error   string
}

func (e RunCommandFailedEvent) SendToDiscord() {
	webhookReq := discordDto.WebhookRequest{}
	content := fmt.Sprintf("Ooppss.. Error while deploying on `%v`\n", e.Config.Appname)
	webhookReq.Content = content

	discord.Send(webhookReq, nil, e.Config.DiscordWebhook)
}
