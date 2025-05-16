package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler func(s *discordgo.Session, m *discordgo.MessageCreate) error

type CommandRouter struct {
	handlers map[string]CommandHandler
}

func NewCommandRouter() *CommandRouter {
	return &CommandRouter{
		handlers: make(map[string]CommandHandler),
	}
}

func (r *CommandRouter) Register(command string, handler CommandHandler) {
	r.handlers[command] = handler
}

func (r *CommandRouter) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := strings.ToLower(strings.TrimSpace(m.Content))

	if handler, ok := r.handlers[cmd]; ok {
		if err := handler(s, m); err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err))
		}
	}
}
