package discord

import (
	log "github.com/sirupsen/logrus"
	"gitlab-webhook/internal/libraries/discord/dto"
	"gitlab-webhook/internal/libraries/http"
)

// Send ..
func Send(dto dto.WebhookRequest, attachment []byte, url string) {
	var err error
	dto.Content = subStr(dto.Content)

	if attachment != nil {
		_, _, err = http.HTTPPostFile(url, dto, attachment)
	} else {
		_, err = http.HTTPPost(url, dto)
	}

	if err != nil {
		log.Println("error while send webhook")
	}
}

// subStr ..
func subStr(content string) string {
	a := []rune(content)
	max := len(a)
	if max <= 2000 {
		return string(a[0:max])
	}
	return string(a[0:1999])
}
