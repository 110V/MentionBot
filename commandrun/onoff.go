package commandrun

import (
	"github.com/110V/MentionBot/users"
	"github.com/bwmarrin/discordgo"
)

func MentionOn(s *discordgo.Session, m *discordgo.MessageCreate, user users.User, args []string) {
	user.Running = true

	err := users.Update(user)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "뭐징??")
		return
	}

	s.ChannelMessageSend(m.ChannelID, ":loud_sound:감지기능이 켜졌습니다!")
}

func MentionOff(s *discordgo.Session, m *discordgo.MessageCreate, user users.User, args []string) {
	user.Running = false

	err := users.Update(user)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "뭐징??")
		return
	}

	s.ChannelMessageSend(m.ChannelID, ":mute:감지기능이 꺼졌습니다.")
}
