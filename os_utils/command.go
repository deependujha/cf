package os_utils

import "os/exec"

// run command: executes a shell command and returns its output.
func RunCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
