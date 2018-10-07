package utils

import (
	"bytes"
	"os/exec"
)

func Exec(command string, arguments ...string) (string, bool) {
	cmd := exec.Command(command, arguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", false
	}

	return out.String(), true
}
