package main

import (
	"strings"
	"sync"

	"github.com/110V/MentionBot/utils"

	"github.com/110V/MentionBot/mention"
	"github.com/110V/MentionBot/userconfig"

	"github.com/110V/MentionBot/commands"
	"github.com/110V/MentionBot/config"
	"github.com/bwmarrin/discordgo"
)

var mutex = &sync.Mutex{}

func newMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//mention.Mention(s, m.Author.ID)
	if m.Author.Bot {
		return
	}
	m.Content = strings.ToLower(m.Content)
	//fmt.Println(m.Content)
	if commands.CheckPerfix(m.Content) && !(m.Content == config.GConfig.Prefix) {
		mutex.Lock()
		err := commands.Run(commands.GetCommandsAndArgs(m.Content), s, m)
		mutex.Unlock()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		}
	} else {
		mutex.Lock()

		if utils.IndexOfString(config.GConfig.ChannelList, m.ChannelID) != -1 {
			for _, user := range userconfig.GUserConfig.Users {
				if !user.Running || user.Id == m.Author.ID {
					continue
				}
				for _, nick := range user.Nicklist {
					if strings.Contains(m.Content, nick) {
						mention.Mention(s, m, user.Id, nick)
						mutex.Unlock()
						return
					}
				}
			}
		}
		mutex.Unlock()

	}
}
