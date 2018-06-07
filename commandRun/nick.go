package commandRun

import (
	"strings"

	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/consts"
	"github.com/110V/MentionBot/userconfig"
	"github.com/bwmarrin/discordgo"
)

func nickMsg(ID string) string {
	u := userconfig.FindUser(&userconfig.GUserConfig, ID)
	if u != nil || len(u.Nicklist) != 0 {
		return ":x: `등록된 닉네임이 없습니다.`"
	} else {
		return ":paperclip:`현재 등록된 닉네임 목록 →`" + strings.Join(userconfig.FindUser(&userconfig.GUserConfig, ID).Nicklist, ",")
	}

}

func NickCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if args == nil {
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
		return
	}
	switch commandtype.CommandMap[args[0]] {
	case commandtype.ADD:
		{
			if len(args) > 1 {
				userconfig.AddNick(m.Author.ID, args[1:])
				s.ChannelMessageSend(m.ChannelID, ":paperclip:`현재 등록된 닉네임 목록 →`"+strings.Join(userconfig.FindUser(&userconfig.GUserConfig, m.Author.ID).Nicklist, ","))
				return
			}
		}
	case commandtype.REMOVE:
		{
			if len(args) > 1 {
				userconfig.RemoveNick(m.Author.ID, args[1:])
				s.ChannelMessageSend(m.ChannelID, ":paperclip:`현재 등록된 닉네임 목록 →`"+strings.Join(userconfig.FindUser(&userconfig.GUserConfig, m.Author.ID).Nicklist, ","))
				return
			}
		}
	case commandtype.RESET:
		{
			userconfig.ResetNick(m.Author.ID)
			s.ChannelMessageSend(m.ChannelID, ":leftwards_arrow_with_hook:`닉네임 목록이 리셋되었습니다.`")
			return
		}
	case commandtype.LIST:
		{
			s.ChannelMessageSend(m.ChannelID, nickMsg(m.Author.ID))
			return
		}
	}
	s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
}
