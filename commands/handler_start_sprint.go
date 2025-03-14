package commands

import (
	//"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandlerStartSprint(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, "!") {
		return
	}

	if m.Content == "/sprint" {
		s.ChannelMessageSend(m.ChannelID, "How many minutes do you want your sprint to last?")
	}
}
