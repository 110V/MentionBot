package commands

import (
	"github.com/110V/MentionBot/commandrun"
	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/users"
	"github.com/bwmarrin/discordgo"
)

var UserCommandList, AdminCommandList []commandtype.TSCommand

func RegistCommands() {
	UserCommandList = []commandtype.TSCommand{
		commandtype.TSCommand{
			Command: commandtype.NICK,
			Use:     "nick add/remove/reset/list",
			Help:    "감지될 별명을 추가/삭제/리셋 합니다. 등록된 닉네임을 보여줍니다.",
			Run:     commandrun.NickCommandHandler,
		},
		commandtype.TSCommand{
			Command: commandtype.ON,
			Use:     "on",
			Help:    "감지 기능을 킵니다. ",
			Run:     commandrun.MentionOn,
		},
		commandtype.TSCommand{
			Command: commandtype.OFF,
			Use:     "off",
			Help:    "감지 기능을 끕니다. ",
			Run:     commandrun.MentionOff,
		},
		commandtype.TSCommand{
			Command: commandtype.HELP,
			Use:     "help (admin)",
			Help:    "도움말을 불러옵니다. ",
			Run:     helpHandler,
		},
	}
	AdminCommandList = []commandtype.TSCommand{
		commandtype.TSCommand{
			Command: commandtype.CHANNEL,
			Use:     "admin channel add/remove/reset #ROOMNAME&here",
			Help:    "감지가 될 채널을 주가/삭제/리셋 합니다. 등록된 채널을 보여줍니다.",
			Run:     commandrun.ChannelCommandHandler,
		},
		commandtype.TSCommand{
			Command: commandtype.RELOAD,
			Use:     "admin reload",
			Help:    "콘피그를 리로드합니다",
			Run:     commandrun.ReloadHandler,
		},
	}
}
func helpHandler(s *discordgo.Session, m *discordgo.MessageCreate, user users.User, args []string) {
	var List []commandtype.TSCommand
	if len(args) != 0 && commandtype.CommandMap[args[0]] == commandtype.ADMIN {
		List = AdminCommandList
	} else {
		List = UserCommandList
	}
	fields := []*discordgo.MessageEmbedField{}
	for _, cmd := range List {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  config.Get().Prefix + cmd.Use,
			Value: ":arrow_right: " + cmd.Help,
		})
	}
	msgEmbed := &discordgo.MessageEmbed{
		Title:  "HELP:question:",
		Fields: fields,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, msgEmbed)
}
