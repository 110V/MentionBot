package commandRun

import (
	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/userconfig"
	"github.com/bwmarrin/discordgo"
)

func ReloadHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	userconfig.OpenConfig()
	config.OpenConfig()
	s.ChannelMessageSend(m.ChannelID, ":open_file_folder: 리로드 완료!")
}
