package config

import (
	"fmt"
	"time"
)

func WebhookOttopayDriver() (*GitlabConfig, error) {
	return &GitlabConfig{
		Appname: "Ottofin/ottopay-driver",
		Token:   "k0d3Rah451a",
		Path:    "/opt/go/src/ottopay-driver",
		Command: []string{
			// Pull from repository
			fmt.Sprintf("%v pull origin development 2>&1", GetGitLocation()),

			// Backup latest build
			fmt.Sprintf("%v ottopay-driver ottopay-driver.%v 2>&1",
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
			fmt.Sprintf("/opt/go/src/ottopay-driver/build-version.sh"),

			// Stop service
			fmt.Sprintf("%v -9 $(%v ax | %v ottopay-driver | %v -v grep | %v '{ print $1 }') 2>&1",
				GetBinaryLocation("kill"),
				GetBinaryLocation("ps"),
				GetBinaryLocation("grep"),
				GetBinaryLocation("fgrep"),
				GetBinaryLocation("awk"),
			),

			// Start service
			fmt.Sprintf("/opt/go/src/ottopay-driver/start.sh 2>&1"),
		},
		User:           GetUser(),
		DiscordWebhook: "https://discord.com/api/webhooks/877485283296297000/YfF5tgT7vg8mWgPXWl0_XBbeecbBd7MY7qdEKI7fCizSw6NcioHphm5gDkDErHHm3RSc",
	}, nil
}
