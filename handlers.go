package main

import (
	"log"
	"strings"

	"github.com/110V/MentionBot/users"
	"github.com/110V/MentionBot/utils"

	"github.com/110V/MentionBot/mention"

	"github.com/110V/MentionBot/commands"
	"github.com/110V/MentionBot/config"
	"github.com/bwmarrin/discordgo"
)

func channelDelete(s *discordgo.Session, c *discordgo.ChannelDelete) {
	conf := config.Get()
	if utils.IndexOfString(conf.ChannelList, c.ID) != -1 {
		for i, ch := range conf.ChannelList {
			if ch == c.ID {
				conf.ChannelList = append(conf.ChannelList[:i], conf.ChannelList[i+1:]...)
				i--
			}
		}
		err := config.Update(conf)
		if err != nil {
			log.Println(err)
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

	if utils.IndexOfString(config.Get().ChannelList, m.ChannelID) != -1 {
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
