package commandRun

import (
	"fmt"
	"regexp"

	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/consts"
	"github.com/110V/MentionBot/utils"
	"github.com/bwmarrin/discordgo"
)

func channelMsg() string {
	if len(config.GConfig.ChannelList) != 0 {
		listStr := ""
		for _, channel := range config.GConfig.ChannelList {
			listStr += "<#" + channel + "> "
		}
		return ":paperclip:`현재 등록된 채널 목록 →`" + listStr
	} else {
		return ":x: `등록된 채널이 없습니다.`"
	}
}

func addChannel(channel string, s *discordgo.Session, chanID string) {
	if utils.IndexOfString(config.GConfig.ChannelList, channel) == -1 {
		config.GConfig.ChannelList = append(config.GConfig.ChannelList, channel)
		config.SaveConfig(config.GConfig)
		s.ChannelMessageSend(chanID, channelMsg())
		return
	}
	s.ChannelMessageSend(chanID, consts.AlreadyExist)
}
func removeChannel(channel string, s *discordgo.Session, chanID string) {
	if utils.IndexOfString(config.GConfig.ChannelList, channel) != -1 {
		tempArr := make([]string, 0)
		for i := range config.GConfig.ChannelList {
			if channel != config.GConfig.ChannelList[i] {
				tempArr = append(tempArr, config.GConfig.ChannelList[i])
			}
		}
		config.GConfig.ChannelList = tempArr
		config.SaveConfig(config.GConfig)
		s.ChannelMessageSend(chanID, channelMsg())
		return
	}
	s.ChannelMessageSend(chanID, consts.NotExist)
}

func ChannelCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if args == nil {
		s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
		return
	}
	switch commandtype.CommandMap[args[0]] {
	case commandtype.ADD:
		{
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
			return
		}
	case commandtype.REMOVE:
		{
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
			return
		}
	case commandtype.RESET:
		{
			config.GConfig.ChannelList = nil
			config.SaveConfig(config.GConfig)
			s.ChannelMessageSend(m.ChannelID, ":leftwards_arrow_with_hook:`채널 목록이 리셋되었습니다.`")
			return
		}
	case commandtype.LIST:
		{
			s.ChannelMessageSend(m.ChannelID, channelMsg())
			return
		}
	}
	s.ChannelMessageSend(m.ChannelID, consts.ArgsError)
}
