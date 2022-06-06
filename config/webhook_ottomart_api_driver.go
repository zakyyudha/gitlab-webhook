package config

import (
	"fmt"
	"time"
)

func WebhookOttomartApiDriver() (*GitlabConfig, error) {
	return &GitlabConfig{
		Appname: "Ottofin/ottomart-api-driver",
		Token:   "k0d3Rah451a",
		Path:    "/opt/go/src/ottomart-api-driver",
		Command: []string{
			// Pull from repository
			fmt.Sprintf("%v pull origin development 2>&1", GetGitLocation()),

			// Backup latest build
			fmt.Sprintf("%v ottomart-api-driver ottomart-api-driver.%v 2>&1",
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
			fmt.Sprintf("/opt/go/src/ottomart-api-driver/build-version.sh"),

			// Stop service
			fmt.Sprintf("%v -9 $(%v ax | %v ottomart-api-driver | %v -v grep | %v '{ print $1 }') 2>&1",
				GetBinaryLocation("kill"),
				GetBinaryLocation("ps"),
				GetBinaryLocation("grep"),
				GetBinaryLocation("fgrep"),
				GetBinaryLocation("awk"),
			),

			// Start service
			fmt.Sprintf("/opt/go/src/ottomart-api-driver/start.sh 2>&1"),
		},
		User:           GetUser(),
		DiscordWebhook: "https://discord.com/api/webhooks/882120158922035210/4q8Ql3N1t9CPId4Rq2hJ0Qd4Y8gpKSADtL409mtQY-OcVdZNelGye-UNs92q6w1rZAHT",
	}, nil
}
