package commandRun

import (
	"github.com/110V/MentionBot/userconfig"
	"github.com/bwmarrin/discordgo"
)

func MentionOn(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	u := userconfig.FindUser(&userconfig.GUserConfig, m.Author.ID)
	if u != nil {
		u.Running = true
	} else {
		userconfig.GUserConfig.Users = append(userconfig.GUserConfig.Users, userconfig.TUser{m.Author.ID, nil, true})
	}
	userconfig.SaveConfig(userconfig.GUserConfig)
	s.ChannelMessageSend(m.ChannelID, ":loud_sound:감지기능이 켜졌습니다!")
}
func MentionOff(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	u := userconfig.FindUser(&userconfig.GUserConfig, m.Author.ID)
	if u != nil {
		u.Running = false
	} else {
		userconfig.GUserConfig.Users = append(userconfig.GUserConfig.Users, userconfig.TUser{m.Author.ID, nil, false})
	}
	userconfig.SaveConfig(userconfig.GUserConfig)
	s.ChannelMessageSend(m.ChannelID, ":mute:감지기능이 꺼졌습니다.")
}
