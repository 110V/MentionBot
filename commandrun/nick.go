package commandrun

import (
	"strings"

	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/consts"
	"github.com/110V/MentionBot/users"
	"github.com/bwmarrin/discordgo"
)

func nickMsg(user users.User) string {
	if len(user.Nicklist) == 0 {
		return ":x: `등록된 닉네임이 없습니다.`"
	}
	return ":paperclip:`현재 등록된 닉네임 목록 →`" + strings.Join(user.Nicklist, ",")
}

func NickCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, user users.User, args []string) {
	if args == nil {
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
		return
	}

	switch commandtype.CommandMap[args[0]] {
	case commandtype.ADD:
		err := user.AddNick(args[1:])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		err = users.Update(user)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "뭐징??")
			return
		}

		s.ChannelMessageSend(m.ChannelID, nickMsg(user))
	case commandtype.REMOVE:
		user.RemoveNick(args[1:])

		err := users.Update(user)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "뭐징??")
			return
		}

		s.ChannelMessageSend(m.ChannelID, nickMsg(user))
	case commandtype.RESET:
		user.ResetNick()

		err := users.Update(user)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "뭐징??")
			return
		}

		s.ChannelMessageSend(m.ChannelID, ":leftwards_arrow_with_hook:`닉네임 목록이 리셋되었습니다.`")
	case commandtype.LIST:
		s.ChannelMessageSend(m.ChannelID, nickMsg(user))
	default:
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
	}

}
