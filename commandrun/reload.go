package commandrun

import (
	"github.com/110V/MentionBot/channels"
	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/users"
	"github.com/bwmarrin/discordgo"
)

func ReloadHandler(s *discordgo.Session, m *discordgo.MessageCreate, user users.User, args []string) {
	err := users.Open()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	err = config.Open()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	err = channels.Open()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	s.ChannelMessageSend(m.ChannelID, ":open_file_folder: 리로드 완료!")
}
