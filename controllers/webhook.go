package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gitlab-webhook/config"
	"gitlab-webhook/dto"
	"gitlab-webhook/events"
	"gitlab-webhook/internal/libraries/discord"
	discordDto "gitlab-webhook/internal/libraries/discord/dto"
	"net/http"
)

// ReceiveWebhook ..
func ReceiveWebhook(c echo.Context) error {
	appname := c.Param("appname")
	appConfig, err := config.Get(appname)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	webHookData := new(dto.GitlabWebhooks)
	if err := c.Bind(webHookData); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if appConfig.Token != c.Request().Header.Get("X-Gitlab-Token") {
		log.Error("suspicious request")
		return c.JSON(http.StatusUnauthorized, "suspicious request")
	}

	go preDeployment(appConfig, webHookData)

	events.Dispatch(events.WebhookReceived,
		events.WebhookReceivedEvent{
			Config:  appConfig,
			Request: webHookData,
		},
	)

	return c.JSON(http.StatusOK, "Webhook received")
}

func preDeployment(gitlabConfig *config.GitlabConfig, webHookData *dto.GitlabWebhooks)  {
	webhookReq := discordDto.WebhookRequest{}
	content := fmt.Sprintf("Starting automation deployment of `%v`...\n", gitlabConfig.Appname)
	webhookReq.Content = content

	var description string
	for _, commit := range webHookData.Commits {
		commitId := string([]rune(commit.ID)[0:8])
		commitUrl := commit.URL
		commitTitle := commit.Title

		description += fmt.Sprintf("[%s](%s): %s \n", commitId, commitUrl, commitTitle)
	}

	webhookReq.Embeds = []discordDto.Embed{
		{
			Description: description,
			Author: struct {
				Name    string `json:"name,omitempty"`
				URL     string `json:"url,omitempty"`
				IconURL string `json:"icon_url,omitempty"`
			}{
				Name:    webHookData.UserName,
				IconURL: webHookData.UserAvatar,
			},
		},
	}

	discord.Send(webhookReq, nil, gitlabConfig.DiscordWebhook)
}