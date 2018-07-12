package commandrun

import (
	"errors"
	"regexp"

	"github.com/110V/MentionBot/channels"
	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/consts"
	"github.com/110V/MentionBot/users"
	"github.com/bwmarrin/discordgo"
)

func channelMsg() string {
	chanlist := channels.Get()
	if len(chanlist) != 0 {
		listStr := ""
		for _, channel := range chanlist {
			listStr += "<#" + channel + "> "
		}
		return ":paperclip:`현재 등록된 채널 목록 →`" + listStr
	} else {
		return ":x: `등록된 채널이 없습니다.`"
	}
}

func transChannelIDArg(arg string, MsgChanID string) (error, string) {

	match, err := regexp.MatchString("<#(?:[0-9]){18}>", arg)
	if err != nil {
		return errors.New(consts.ArgsError), ""
	}
	if match {
		return nil, arg[2:20]
	}
	if commandtype.CommandMap[arg] == commandtype.HERE {
		return nil, MsgChanID
	}
	return errors.New(consts.ArgsError), ""
}

func ChannelCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, user users.User, args []string) {
	if args == nil {
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
		return
	}
	switch commandtype.CommandMap[args[0]] {
	case commandtype.ADD:
		if len(args) > 1 {
			err, ID := transChannelIDArg(args[1], m.ChannelID)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, err.Error())
			}
			channels.AddChannel(ID)
			s.ChannelMessageSend(m.ChannelID, channelMsg())
			return
		}
	case commandtype.REMOVE:
		if len(args) > 1 {
			err, ID := transChannelIDArg(args[1], m.ChannelID)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, err.Error())
			}
			channels.AddChannel(ID)
			s.ChannelMessageSend(m.ChannelID, channelMsg())
			return
		}
	case commandtype.RESET:
		err := channels.Reset()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, consts.InternalError)
			return
		}
		s.ChannelMessageSend(m.ChannelID, ":leftwards_arrow_with_hook:`채널 목록이 리셋되었습니다.`")
		return
	case commandtype.LIST:
		s.ChannelMessageSend(m.ChannelID, channelMsg())
		return
	}
	s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
}
