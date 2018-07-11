package commands

import (
	"errors"
	"log"
	"strings"

	"github.com/110V/MentionBot/commandtype"
	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/consts"
	"github.com/110V/MentionBot/users"
	"github.com/bwmarrin/discordgo"
)

func CheckPerfix(command string) bool {
	return strings.HasPrefix(command, config.Get().Prefix) && command != config.Get().Prefix
}

func GetCommandsAndArgs(command string) []string {
	return strings.Fields(strings.TrimPrefix(command, config.Get().Prefix))
}

func FindCommand(cmdType commandtype.Command, commandList []commandtype.TSCommand) commandtype.TSCommand {
	for _, tsc := range commandList {
		if tsc.Command == cmdType {
			return tsc
		}
	}
	return commandtype.TSCommand{Command: 0}
}

func Run(commandAndArgs []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	var tsc commandtype.TSCommand

	if commandtype.CommandMap[commandAndArgs[0]] == commandtype.ADMIN {
		if len(commandAndArgs) < 2 {
			return errors.New(consts.ArgsError)
		}

		tsc = FindCommand(commandtype.CommandMap[commandAndArgs[1]], AdminCommandList)
		if tsc.Command == 0 {
			return errors.New(consts.InvalidCommand)
		}

		if !config.CheckAdmin(m.Author.ID) {
			return errors.New(consts.PermissionDenied)
		}

		commandAndArgs = commandAndArgs[1:]
	} else {
		tsc = FindCommand(commandtype.CommandMap[commandAndArgs[0]], UserCommandList)
		if tsc.Command == 0 {
			return errors.New(consts.InvalidCommand)
		}
	}

	user, err := users.GetOrCreate(m.Author.ID)
	if err != nil {
		//DDUTTA
		log.Println(err)
		return errors.New(consts.InternalError)
	}

	if len(commandAndArgs) > 1 {
		tsc.Run(s, m, user, commandAndArgs[1:])
	} else {
		tsc.Run(s, m, user, nil)
	}

	return nil
}
