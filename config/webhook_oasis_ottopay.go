package config

import (
	"fmt"
	"time"
)

func WebhookOasisOttopay() (*GitlabConfig, error) {
	return &GitlabConfig{
		Appname: "Ottopay/oasis-ottopay",
		Token:   "k0d3Rah451a",
		Path:    "/opt/go/src/oasis-ottopay",
		Command: []string{
			// Pull from repository
			fmt.Sprintf("%v pull origin development 2>&1", GetGitLocation()),

			// Backup latest build
			fmt.Sprintf("%v oasis-ottopay oasis-ottopay.%v 2>&1",
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
			fmt.Sprintf("/opt/go/src/oasis-ottopay/build-version.sh"),

			// Stop service
			fmt.Sprintf("%v -9 $(%v ax | %v oasis-ottopay | %v -v grep | %v '{ print $1 }') 2>&1",
				GetBinaryLocation("kill"),
				GetBinaryLocation("ps"),
				GetBinaryLocation("grep"),
				GetBinaryLocation("fgrep"),
				GetBinaryLocation("awk"),
			),

			// Start service
			fmt.Sprintf("/opt/go/src/oasis-ottopay/start.sh 2>&1"),
		},
		User:           GetUser(),
		DiscordWebhook: "https://discord.com/api/webhooks/877918990863839242/mwLPaOck3fzFmDBn8kA05wJRmrqVTVVKsa_bDseNLonW58Ey8nwJbMXKDi-KWCDuewPe",
	}, nil
}
