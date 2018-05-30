package commands

import (
	"github.com/110V/MentionBot/CommandRun"
	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/config"
)

var UserCommandList []commandtype.TSCommand = []commandtype.TSCommand{
	commandtype.TSCommand{
		Command: commandtype.NICK,
		Help:    ":question:" + config.GConfig.Prefix + "nick add/remove/reset :arrow_right:  감지될 별명을 추가/삭제/리셋 합니다.",
		Run:     CommandRun.NickCommandHandler},
}
var AdminCommandList []commandtype.TSCommand = []commandtype.TSCommand{
	commandtype.TSCommand{
		Command: commandtype.CHANNEL,
		Help:    ":question:" + config.GConfig.Prefix + "admin channel add/remove/reset #ROOMNAME&here :arrow_right:  감지가 될 채널을 주가/삭제/리셋 합니다.",
		Run:     nil},
}
