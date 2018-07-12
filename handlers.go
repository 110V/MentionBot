package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/110V/MentionBot/channels"
	"github.com/110V/MentionBot/users"
	"github.com/110V/MentionBot/utils"

	"github.com/110V/MentionBot/mention"

	"github.com/110V/MentionBot/commands"
	"github.com/bwmarrin/discordgo"
)

var mu sync.Mutex

func channelDelete(s *discordgo.Session, c *discordgo.ChannelDelete) {
	mu.Lock()
	defer mu.Unlock()
	if utils.IndexOfString(channels.Get(), c.ID) != -1 {
		err := channels.RemoveChannel(c.ID)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}

func newMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	m.Content = strings.ToLower(m.Content)

	if commands.CheckPerfix(m.Content) {
		err := commands.Run(commands.GetCommandsAndArgs(m.Content), s, m)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		}
		return
	}

	if utils.IndexOfString(channels.Get(), m.ChannelID) != -1 {
		for _, user := range users.GetAll() {
			if !user.Running || user.ID == m.Author.ID {
				continue
			}
			for _, nick := range user.Nicklist {
				if strings.Contains(m.Content, nick) {
					mention.Mention(s, m, user.ID, nick)
					break
				}
			}
		}
	}
}
