package config

import (
	"fmt"
	"time"
)

func WebhookOttopayEmoneyDriver() (*GitlabConfig, error) {
	return &GitlabConfig{
		Appname: "Ottofin/ottopay-e-money-driver",
		Token:   "k0d3Rah451a",
		Path:    "/opt/go/src/ottopay-e-money-driver",
		Command: []string{
			// Pull from repository
			fmt.Sprintf("%v pull origin development 2>&1", GetGitLocation()),

			// Backup latest build
			fmt.Sprintf("%v ottopay-e-money-driver ottopay-e-money-driver.%v 2>&1",
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
			fmt.Sprintf("/opt/go/src/ottopay-e-money-driver/build-version.sh"),

			// Stop service
			fmt.Sprintf("%v -9 $(%v ax | %v ottopay-e-money-driver | %v -v grep | %v '{ print $1 }') 2>&1",
				GetBinaryLocation("kill"),
				GetBinaryLocation("ps"),
				GetBinaryLocation("grep"),
				GetBinaryLocation("fgrep"),
				GetBinaryLocation("awk"),
			),

			// Start service
			fmt.Sprintf("/opt/go/src/ottopay-e-money-driver/start.sh 2>&1"),
		},
		User:           GetUser(),
		DiscordWebhook: "https://discord.com/api/webhooks/882126673494036491/yxCcfXkWbu1mohHYET8R06ghsgu88Nei5nVw3x38rglIxmXqAGpE6I4j2MCK_srU3jJr",
	}, nil
}
