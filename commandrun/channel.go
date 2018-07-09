package commandrun

import (
	"fmt"
	"regexp"

	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/consts"
	"github.com/110V/MentionBot/users"
	"github.com/110V/MentionBot/utils"
	"github.com/bwmarrin/discordgo"
)

func channelMsg(conf config.Config) string {
	if len(conf.ChannelList) != 0 {
		listStr := ""
		for _, channel := range conf.ChannelList {
			listStr += "<#" + channel + "> "
		}
		return ":paperclip:`현재 등록된 채널 목록 →`" + listStr
	} else {
		return ":x: `등록된 채널이 없습니다.`"
	}
}

func addChannel(channel string, s *discordgo.Session, chanID string) {
	conf := config.Get()
	if utils.IndexOfString(conf.ChannelList, channel) != -1 {
		s.ChannelMessageSend(chanID, consts.AlreadyExist)
		return
	}

	conf.ChannelList = append(conf.ChannelList, channel)
	err := config.Update(conf)
	if err != nil {
		s.ChannelMessageSend(chanID, "뭐징??")
		return
	}

	s.ChannelMessageSend(chanID, channelMsg(conf))
}

func removeChannel(channel string, s *discordgo.Session, chanID string) {
	conf := config.Get()
	if utils.IndexOfString(conf.ChannelList, channel) != -1 {
		tempArr := make([]string, 0)
		for i := range conf.ChannelList {
			if channel != conf.ChannelList[i] {
				tempArr = append(tempArr, conf.ChannelList[i])
			}
		}
		conf.ChannelList = tempArr
		err := config.Update(conf)
		if err != nil {
			s.ChannelMessageSend(chanID, "뭐징??")
			return
		}

		s.ChannelMessageSend(chanID, channelMsg(conf))
		return
	}
	s.ChannelMessageSend(chanID, consts.NotExist)
}

func ChannelCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, user users.User, args []string) {
	if args == nil {
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
		return
	}
	switch commandtype.CommandMap[args[0]] {
	case commandtype.ADD:
		if len(args) > 1 {

			match, err := regexp.MatchString("<#(?:[0-9]){18}>", args[1])
			if err != nil {
				fmt.Println(err.Error())
				s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
				return
			}
			if match {
				addChannel(args[1][2:20], s, m.ChannelID)
				return
			} else if commandtype.CommandMap[args[1]] == commandtype.HERE {
				addChannel(m.ChannelID, s, m.ChannelID)
				return
			}
		}
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
	case commandtype.REMOVE:
		if len(args) > 1 {
			match, err := regexp.MatchString("<#(?:[0-9]){18}>", args[1])
			if err != nil {
				fmt.Println(err.Error())
				s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
				return
			}
			if match {
				removeChannel(args[1][2:20], s, m.ChannelID)
				return
			} else if commandtype.CommandMap[args[1]] == commandtype.HERE {
				removeChannel(m.ChannelID, s, m.ChannelID)
				return
			} else {
				s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
				return
			}

		}
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
	case commandtype.RESET:
		conf := config.Get()
		conf.ChannelList = nil
		err := config.Update(conf)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "뭐징??")
		}
		s.ChannelMessageSend(m.ChannelID, ":leftwards_arrow_with_hook:`채널 목록이 리셋되었습니다.`")
	case commandtype.LIST:
		s.ChannelMessageSend(m.ChannelID, channelMsg(config.Get()))
	default:
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
	}
}
