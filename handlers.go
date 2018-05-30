package main

import (
	"strings"

	"github.com/110V/MentionBot/commands"
	"github.com/110V/MentionBot/consts"
	"github.com/bwmarrin/discordgo"
)

func NewMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	m.Content = strings.ToLower(m.Content)
	if commands.CheckPerfix(m.Content) && !m.Author.Bot {
		if !commands.Run(commands.GetCommandsAndArgs(m.Content), *s, *m) {
			s.ChannelMessageSend(m.ChannelID, consts.InvalidCommand)
		}
	}
}
