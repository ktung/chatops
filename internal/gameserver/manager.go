package gameserver

import (
	"fmt"
	"os/exec"
)

type ServerManager struct {
	cmdPath string
}

func NewServerManager(cmdPath string) *ServerManager {
	return &ServerManager{
		cmdPath: cmdPath,
	}
}

func (m *ServerManager) ExecuteCommand(serverType, action string) (string, error) {
	scriptName := fmt.Sprintf("%s_%s.ps1", serverType, action)
	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command", m.cmdPath+scriptName)

	if action == "info" {
		output, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("failed to execute %s: %w", scriptName, err)
		}
		return string(output), nil
	}

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute %s: %w", scriptName, err)
	}
	return fmt.Sprintf("%s %s completed successfully", serverType, action), nil
}
