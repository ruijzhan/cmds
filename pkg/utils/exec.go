package utils

import "os/exec"

func ExecuteCommand(cmd, subcmd string, args ...string) (string, error) {
	args = append([]string{subcmd}, args...)
	c := exec.Command(cmd, args...)
	bytes, err := c.CombinedOutput()
	return string(bytes), err
}
