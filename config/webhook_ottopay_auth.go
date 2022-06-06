package config

import (
	"fmt"
	"time"
)

func WebhookOttopayAuth() (*GitlabConfig, error) {
	return &GitlabConfig{
		Appname: "Ottopay/ottopay-auth-services",
		Token:   "k0d3Rah451a",
		Path:    "/opt/go/src/ottopay-auth-services",
		Command: []string{
			// Pull from repository
			fmt.Sprintf("%v pull origin development 2>&1", GetGitLocation()),

			// Backup latest build
			fmt.Sprintf("%v ottopay-auth-services ottopay-auth-services.%v 2>&1",
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
			fmt.Sprintf("/opt/go/src/ottopay-auth-services/build-version.sh"),

			// Stop service
			fmt.Sprintf("%v -9 $(%v ax | %v ottopay-auth-services | %v -v grep | %v '{ print $1 }') 2>&1",
				GetBinaryLocation("kill"),
				GetBinaryLocation("ps"),
				GetBinaryLocation("grep"),
				GetBinaryLocation("fgrep"),
				GetBinaryLocation("awk"),
			),

			// Start service
			fmt.Sprintf("/opt/go/src/ottopay-auth-services/start.sh 2>&1"),
		},
		User:           GetUser(),
		DiscordWebhook: "https://discord.com/api/webhooks/877864010878451723/DXc7FloUoTF9d9pudV5b5bNqYiYVyvzQQr6WN1XFv48KaG9Uc1x5RD0m1eO3vtqFlnv1",
	}, nil
}
