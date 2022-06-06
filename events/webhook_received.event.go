package events

import (
	"bytes"
	"fmt"
	"github.com/acarl005/stripansi"
	log "github.com/sirupsen/logrus"
	"gitlab-webhook/config"
	"gitlab-webhook/dto"
	"gitlab-webhook/internal/libraries/discord"
	discordDto "gitlab-webhook/internal/libraries/discord/dto"
	"gitlab-webhook/internal/libraries/event"
	"gitlab-webhook/utils"
	"os"
	"time"
)

// WebhookReceived ..
const WebhookReceived event.Name = "WebhookReceived"

// WebhookReceivedEvent ..
type WebhookReceivedEvent struct {
	Config  *config.GitlabConfig
	Request *dto.GitlabWebhooks
}

// RunCommand ..
func (e WebhookReceivedEvent) RunCommand() {
	var results []string
	var preResults string
	var errGlobal error
	deploymentSucceeded := true

	for _, cmd := range e.Config.Command {

		errGlobal = os.Chdir(e.Config.Path)
		if errGlobal != nil {
			deploymentSucceeded = false
			log.Error("Error while changing directory")
			go dispatcher.Dispatch(RunCommandFailed,
				RunCommandFailedEvent{
					Config:  e.Config,
					Request: e.Request,
					Error:   fmt.Sprintf("Error while changing directory: `%v`", errGlobal.Error()),
				},
			)
			return
		}

		cmdResult, errGlobal := utils.ExecCommand(cmd)
		preResults = "Run Command => " + cmd + "\n" + cmdResult + "\n\n"

		log.WithFields(log.Fields{
			"appname": e.Config.Appname,
			"at":      time.Now().Format("2006-01-02 15:04:05"),
			"user":    e.Config.User,
			"cmd":     cmd,
			"result":  cmdResult,
		})

		results = append(results, preResults)

		if errGlobal != nil {
			deploymentSucceeded = false
			go dispatcher.Dispatch(RunCommandFailed,
				RunCommandFailedEvent{
					Config:  e.Config,
					Request: e.Request,
					Error:   fmt.Sprintf("Error while executing command: `%v`", cmd),
				},
			)
			break
		}
	}

	if deploymentSucceeded {
		go dispatcher.Dispatch(RunCommandSucceeded,
			RunCommandSucceededEvent{
				Config:  e.Config,
				Request: e.Request,
			},
		)
	}

	webhookReq := discordDto.WebhookRequest{}
	strResult := mapToStr(results)
	content := fmt.Sprintf("Result of deployment `%v`\n", e.Config.Appname)
	webhookReq.Content = content

	discord.Send(webhookReq, []byte(strResult), e.Config.DiscordWebhook)
}

func mapToStr(m []string) string {
	b := new(bytes.Buffer)
	for _, value := range m {
		_, _ = fmt.Fprintf(b, "%s\n", value)
	}
	return stripansi.Strip(b.String())
}
