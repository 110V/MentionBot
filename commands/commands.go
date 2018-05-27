package commands

import (
	"strings"

	setting "github.com/110V/MentionBot/config"
)

func CheckPerfix(command string) bool {
	return strings.HasPrefix(command, setting.Prefix)
}

func GetCommandsAndArgs(command string) []string {
	return strings.Fields(strings.TrimPrefix(command, setting.Prefix))
}

func Run(CommandAndArgs []string) {
	switch CommandMap[CommandAndArgs[0]] {
	case SET:
		{
			switch CommandMap[CommandAndArgs[1]] {

			}
		}
	case ADD:
		{
			switch CommandMap[CommandAndArgs[1]] {

			}
		}
	case ADMIN:
		{
			switch CommandMap[CommandAndArgs[1]] {

			}
		}
	default:
		{

		}
	}
}
