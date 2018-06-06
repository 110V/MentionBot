package commands

import (
	"errors"
	"strings"

	"github.com/110V/MentionBot/utils"

	"github.com/110V/MentionBot/consts"

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

func Run(commandAndArgs []string, s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	tsc := FindCommand(commandtype.CommandMap[commandAndArgs[0]], UserCommandList)

	if tsc.Command == 0 {
		if commandtype.CommandMap[commandAndArgs[0]] == commandtype.ADMIN {
			if len(commandAndArgs) < 2 {
				err = errors.New(consts.ArgsError)
				return
			}
			tsc = FindCommand(commandtype.CommandMap[commandAndArgs[1]], AdminCommandList)
			if tsc.Command == 0 {
				err = errors.New(consts.InvalidCommand)
				return
			}
			if !utils.CheckAdmin(m.Author.ID) {
				err = errors.New(consts.PermissionDenied)
				return
			}
			commandAndArgs = commandAndArgs[1:]
		} else {
			err = errors.New(consts.InvalidCommand)
			return
		}
	}

	if len(commandAndArgs) > 1 {
		tsc.Run(s, m, commandAndArgs[1:])
	} else {
		tsc.Run(s, m, nil)
	}

	return
}
