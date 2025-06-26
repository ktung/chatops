package cmds

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler func(s *discordgo.Session, m *discordgo.MessageCreate) error

type CommandsMap struct {
	handlers map[string]CommandHandler
}

func NewRouterCommands() *CommandsMap {
	router := &CommandsMap{
		handlers: map[string]CommandHandler{
			"work":               WorkCommand,
			"break":              BreakCommand,
			"enshrouded_info":    EnshroudedInfoCommand,
			"enshrouded_restart": EnshroudedRestartCommand,
			"enshrouded_update":  EnshroudedUpdateCommand,
			"palserver_info":     PalworldInfoCommand,
			"palserver_restart":  PalworldRestartCommand,
		},
	}

	router.handlers["help"] = NewHelpCommand(router)
	return router
}

func (c *CommandsMap) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	messageFields := strings.Fields(m.Content)
	if len(messageFields) == 0 {
		return
	}

	cmd := messageFields[0]
	if handler, ok := c.handlers[cmd]; ok {
		if err := handler(s, m); err != nil {
			log.Printf("Error handling command: %v", err)
		}
	}
}

func (c *CommandsMap) GetAllCommands() []string {
	commands := make([]string, 0, len(c.handlers))
	for command := range c.handlers {
		commands = append(commands, command)
	}
	return commands
}
