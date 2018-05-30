package commands

import (
	"strings"

	"github.com/110V/MentionBot/commandtype"

	"github.com/110V/MentionBot/config"
	"github.com/bwmarrin/discordgo"
)

func CheckPerfix(command string) bool {
	return strings.HasPrefix(command, config.GConfig.Prefix)
}

func GetCommandsAndArgs(command string) []string {
	return strings.Fields(strings.TrimPrefix(command, config.GConfig.Prefix))
}

func FindCommand(cmdType commandtype.Command, commandList []commandtype.TSCommand) commandtype.TSCommand {
	for _, tsc := range commandList {
		if tsc.Command == cmdType {
			return tsc
		}
	}
	return commandtype.TSCommand{Command: 0}
}

func Run(CommandAndArgs []string, s discordgo.Session, m discordgo.MessageCreate) bool {
	tsc := FindCommand(commandtype.CommandMap[CommandAndArgs[0]], UserCommandList)
	if tsc.Command == 0 {
		tsc = FindCommand(commandtype.CommandMap[CommandAndArgs[0]], AdminCommandList)
		if tsc.Command == 0 {
			return false
		}
	}
	if len(CommandAndArgs) > 1 {
		tsc.Run(s, m, CommandAndArgs[1:])
	} else {
		tsc.Run(s, m, nil)
	}

	return true
}
