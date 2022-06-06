package config

import (
	"fmt"
	"time"
)

func WebhookOttopayNotif() (*GitlabConfig, error) {
	return &GitlabConfig{
		Appname: "Ottopay/ottopay-notif",
		Token:   "k0d3Rah451a",
		Path:    "/opt/go/src/ottopay-notif",
		Command: []string{
			// Pull from repository
			fmt.Sprintf("%v pull origin development 2>&1", GetGitLocation()),

			// Backup latest build
			fmt.Sprintf("%v ottopay-notif ottopay-notif.%v 2>&1",
				GetBinaryLocation("cp"),
				time.Now().Format("02-Jan-2006_15:04:05"),
			),

			// Get package from go.sum
			fmt.Sprintf("%v get -v 2>&1",
				GetBinaryLocation("go"),
			),

			// Build latest
			//fmt.Sprintf("%v build -v 2>&1",
			//	GetBinaryLocation("go"),
			//),

			// Build version
			fmt.Sprintf("/opt/go/src/ottopay-notif/build-version.sh"),

			// Stop service
			fmt.Sprintf("%v -9 $(%v ax | %v ottopay-notif | %v -v grep | %v '{ print $1 }') 2>&1",
				GetBinaryLocation("kill"),
				GetBinaryLocation("ps"),
				GetBinaryLocation("grep"),
				GetBinaryLocation("fgrep"),
				GetBinaryLocation("awk"),
			),

			// Start service
			fmt.Sprintf("/opt/go/src/ottopay-notif/start_gin.sh 2>&1"),
		},
		User:           GetUser(),
		DiscordWebhook: "https://discord.com/api/webhooks/882126488747524117/Mum5x_WPjhnY7RKz90IM-YahdlX9GHDnM9TGxKTg5WvOGfrLQM2MramZNu8wJf0YnYhR",
	}, nil
}
