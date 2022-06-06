package config

import (
	"fmt"
	"time"
)

func WebhookOttoverify() (*GitlabConfig, error) {
	return &GitlabConfig{
		Appname: "Ottopay/ottoverify",
		Token:   "k0d3Rah451a",
		Path:    "/opt/go/src/ottoverify",
		Command: []string{
			// Pull from repository
			fmt.Sprintf("%v pull origin development 2>&1", GetGitLocation()),

			// Backup latest build
			fmt.Sprintf("%v ottoverify ottoverify.%v 2>&1",
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
			fmt.Sprintf("/opt/go/src/ottoverify/build-version.sh"),

			// Stop service
			fmt.Sprintf("%v -9 $(%v ax | %v ottoverify | %v -v grep | %v '{ print $1 }') 2>&1",
				GetBinaryLocation("kill"),
				GetBinaryLocation("ps"),
				GetBinaryLocation("grep"),
				GetBinaryLocation("fgrep"),
				GetBinaryLocation("awk"),
			),

			// Start service
			fmt.Sprintf("/opt/go/src/ottoverify/start.sh 2>&1"),
		},
		User:           GetUser(),
		DiscordWebhook: "https://discord.com/api/webhooks/890147899567140905/ZAJ9ltmSyq0SV58Ex0YCprN9E7YG2mgyqv0EM7E2Z2lnBKmSujwx8EndcaxppGOnQ-2A",
	}, nil
}
