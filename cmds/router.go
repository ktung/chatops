package cmds

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler func(s *discordgo.Session, m *discordgo.MessageCreate) error

type CommandsMap struct {
	handlers map[string]CommandHandler
}

func NewRouterCommands() *CommandsMap {
	router := &CommandsMap{
		handlers: map[string]CommandHandler{
			"ping":               PingCommand,
			"pong":               PongCommand,
			"enshrouded_info":    EnshroudedInfoCommand,
			"enshrouded_restart": EnshroudedRestartCommand,
			"palserver_info":     PalworldInfoCommand,
			"palserver_restart":  PalworldRestartCommand,
		},
	}

	router.handlers["help"] = NewHelpCommand(router)
	return router
}

func (c *CommandsMap) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if handler, ok := c.handlers[m.Content]; ok {
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
