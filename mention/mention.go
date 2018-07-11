package mention

import (
	"github.com/bwmarrin/discordgo"
)

func Mention(s *discordgo.Session, m *discordgo.MessageCreate, ID string, nick string) {

	channel, err := s.UserChannelCreate(ID)
	if err == nil {
		Channel, _ := s.Channel(m.ChannelID)
		s.ChannelMessageSend(channel.ID, "<@"+ID+">님, 채널<#"+m.ChannelID+">에서 '"+nick+"'이(가) 감지되었습니다.\n바로가기:arrow_right:https://discordapp.com/channels/"+
			Channel.GuildID+"/"+m.ChannelID+"/"+m.ID)
	}
}
