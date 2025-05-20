package internal

import "os/exec"

func ExecutePowerShellCommand(command string) (string, error) {
	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func RunPowerShellCommand(command string) (error) {
	err := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command", command).Run()
	return err
}
