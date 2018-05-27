package commands

import (
	"strings"
)

func CheckPerfix(command string) bool {
	return strings.HasPrefix(command, setting.Prefix)
}

func GetCommandsAndArgs(commnad string) []string {
	return strings.Fields(strings.TrimPrefix(commnad, setting.Prefix))
}

func Run(CommandAndArgs []string) {

}

//func CheckCommand
