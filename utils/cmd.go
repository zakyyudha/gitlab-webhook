package utils

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
)

// ExecCommand ..
func ExecCommand(cmd string) (string, error) {
	log.Println("Runing Command => " + cmd + " ...\n")
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Println("error while execute command ===>> ["+cmd+"]")
		return string(out), err
	}
	return string(out), nil
}
