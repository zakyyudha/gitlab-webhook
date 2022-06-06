package jobs

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gitlab-webhook/config"
	"gitlab-webhook/internal/libraries/discord"
	"gitlab-webhook/internal/libraries/discord/dto"
)

// Segame ..
func Segame() {
	var webhookReq dto.WebhookRequest
	data := []byte(`{"content":"🔊 Notification","embeds":[{"description":"Se-game bisa kalih 😄","url":"https://discordapp.com","color":9785224,"thumbnail":{"url":"https://media.giphy.com/media/eHWQjrSpsAoMkQPWOV/giphy.gif"},"author":{"name":"Guysss!!!","url":"https://discord.gg/WWqDEjA4","icon_url":"https://cdn.discordapp.com/embed/avatars/0.png"}}]}`)
	err := json.Unmarshal(data, &webhookReq)
	if err != nil {
		log.Println(err.Error())
		return
	}
	conf, err := config.Get("ottopay-loan")
	if err != nil {
		log.Println(err.Error())
		return
	}
	discord.Send(webhookReq, []byte(``), conf.DiscordWebhook)
}
