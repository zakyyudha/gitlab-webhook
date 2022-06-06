package config

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

// GitlabConfig ..
type GitlabConfig struct {
	Appname        string   `json:"appname" example:"ottopay-loan"`
	Token          string   `json:"token" example:"matching-token-with-gitlab"`
	Path           string   `json:"path" example:"/your/dir/to/project"`
	User           string   `json:"user" example:"centos"`
	Command        []string `json:"command" example:"git pull origin development"`
	DiscordWebhook string   `json:"discord_webhook"`
}

// GetGitLocation ..
func GetGitLocation() string {
	out, err := exec.Command("bash", "-c", "which git").Output()
	if err != nil {
		log.Println("git not found => ", err)
	}
	return strings.TrimSpace(string(out))
}

// GetDockerLocation ..
func GetDockerLocation() string {
	out, err := exec.Command("bash", "-c", "which docker").Output()
	if err != nil {
		log.Println("docker not found")
	}
	return strings.TrimSpace(string(out))
}

// GetDockerComposeLocation ..
func GetDockerComposeLocation() string {
	out, err := exec.Command("bash", "-c", "which docker-compose").Output()
	if err != nil {
		log.Println("docker not found")
	}
	return strings.TrimSpace(string(out))
}

// GetShLocation ..
func GetShLocation() string {
	out, err := exec.Command("bash", "-c", "which sh").Output()
	if err != nil {
		log.Println("sh not found")
	}
	return strings.TrimSpace(string(out))
}

// GetUser ..
func GetUser() string {
	out, err := exec.Command("bash", "-c", "whoami").Output()
	if err != nil {
		log.Println("unknown error")
	}
	return strings.TrimSpace(string(out))
}

// GetBinaryLocation ..
func GetBinaryLocation(name string) string {
	out, err := exec.Command("bash", "-c", "which "+name).Output()
	if err != nil {
		log.Println("binary " + name + " not found")
	}
	return strings.TrimSpace(string(out))
}
