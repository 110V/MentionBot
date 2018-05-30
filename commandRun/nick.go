package CommandRun

import (
	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/consts"
	"github.com/110V/MentionBot/nicks"
	"github.com/bwmarrin/discordgo"
)

func NickCommandHandler(s discordgo.Session, m discordgo.MessageCreate, args []string) {
	if args == nil {
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
		return
	}
	switch commandtype.CommandMap[args[0]] {
	case commandtype.ADD:
		{
			if len(args) > 1 {
				nicks.AddNick(m.Author.ID, args[1:])
				s.ChannelMessageSend(m.ChannelID, ":pencil2:`입력하신 닉네임들이 목록에 추가되었습니다.`")
				return
			}
		}
	case commandtype.REMOVE:
		{
			if len(args) > 1 {
				nicks.RemoveNick(m.Author.ID, args[1:])
				s.ChannelMessageSend(m.ChannelID, ":x: `입력하신 닉네임들이 목록에서 제거되었습니다.`")
				return
			}
		}
	case commandtype.RESET:
		{
			nicks.ResetNick(m.Author.ID)
			s.ChannelMessageSend(m.ChannelID, ":leftwards_arrow_with_hook:`닉네임 목록이 리셋되었습니다.`")
			return
		}
	}
	s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
}
